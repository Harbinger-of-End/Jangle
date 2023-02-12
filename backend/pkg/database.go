package pkg

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _db Database = new(database)

type Database interface {
	InitializeDB(ctx context.Context) (postgresErr, mongoErr error)
	UsernameExists(ctx context.Context, username string) bool
	EmailExists(ctx context.Context, email string) bool
	AddUser(
		ctx context.Context,
		userId string,
		username string,
		email string,
		password string,
		tokenType string,
		accessToken string,
		refreshToken string,
		expiresIn time.Duration,
		scope Scope,
	) error
	GetUserId(ctx context.Context, email string) (string, error)
	Close(ctx context.Context) (postgresErr, mongoErr error)
}

type database struct {
	postgresConn *pgx.Conn
	mongoClient  *mongo.Client
}

func (db *database) InitializeDB(
	ctx context.Context,
) (postgresErr, mongoErr error) {
	mongoUri := os.Getenv("MONGODB_URI")
	postgresUri := os.Getenv("POSTGRES_URI")

	db.mongoClient, mongoErr = mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoUri),
	)

	db.postgresConn, postgresErr = pgx.Connect(
		ctx,
		postgresUri,
	)

	return postgresErr, mongoErr
}

func (db database) UsernameExists(
	ctx context.Context,
	username string,
) bool {
	var expected int32
	err := db.postgresConn.QueryRow(
		ctx,
		"SELECT 1 FROM users WHERE Username = $1",
		username,
	).Scan(&expected)

	return err == nil && expected == 1
}

func (db database) EmailExists(
	ctx context.Context,
	email string,
) bool {
	var expected int32
	err := db.postgresConn.QueryRow(
		ctx,
		"SELECT 1 FROM users WHERE Email = $1",
		email,
	).Scan(&expected)

	return err == nil && expected == 1
}

func (db database) AddUser(
	ctx context.Context,
	userId string,
	username string,
	email string,
	password string,
	tokenType string,
	accessToken string,
	refreshToken string,
	expiresIn time.Duration,
	scope Scope,
) error {
	var err error
	_, err = db.postgresConn.Exec(
		ctx,
		"INSERT INTO users VALUES ($1, $2, $3, $4, $5)",
		userId,
		username,
		email,
		password,
		false,
	)

	expiresAt := time.Now().Add(expiresIn * time.Second).UTC()

	if err == nil {
		err = db.AddTokenPair(
			ctx,
			tokenType,
			accessToken,
			refreshToken,
			expiresAt,
			userId,
			scope,
			nil,
		)
	}

	return err
}

func (db database) AddTokenPair(
	ctx context.Context,
	tokenType string,
	accessToken string,
	refreshToken string,
	expiresAt time.Time,
	userId string,
	scope Scope,
	clientId any,
) error {
	_, err := db.postgresConn.Exec(
		ctx,
		"INSERT INTO token_pairs VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		uuid.NewString(),
		tokenType,
		accessToken,
		refreshToken,
		expiresAt,
		userId,
		scope,
		clientId,
	)
	return err
}

func (db database) GetUserId(
	ctx context.Context,
	email string,
) (string, error) {
	var userId string = ""
	err := db.postgresConn.QueryRow(
		ctx,
		"SELECT UserID FROM users WHERE Email = $1",
		email,
	).Scan(&userId)

	return userId, err
}

func (db *database) Close(
	ctx context.Context,
) (postgresErr, mongoErr error) {
	postgresErr = db.postgresConn.Close(ctx)
	mongoErr = db.mongoClient.Disconnect(ctx)

	return postgresErr, mongoErr
}

func Db() Database {
	return _db
}
