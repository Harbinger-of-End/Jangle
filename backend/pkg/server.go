package pkg

import (
	"context"
	"jangle/backend/auth"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var wg sync.WaitGroup

type Server struct {
	auth.UnimplementedAuthenticationServer
}

func (Server) Signup(
	ctx context.Context,
	request *auth.SignupRequest,
) (*auth.SignupResponse, error) {
	start := time.Now().UnixMilli()
	username := request.GetUsername()
	email := request.GetEmail()
	password := request.GetPassword()

	if username == "" || email == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "One of the fields is empty")
	}

	if Db().UsernameExists(ctx, username) {
		return nil, status.Error(codes.AlreadyExists, "User with this username already exists")
	}

	if Db().EmailExists(ctx, email) {
		return nil, status.Error(codes.AlreadyExists, "User with this email already exists")
	}

	userId := GenerateSnowflake()
	accessToken := GenerateToken(userId)
	refreshToken := GenerateToken(userId)
	expiresIn := 60 * 60 * 24 * 60 * time.Second

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Db().AddUser(
			context.Background(),
			userId,
			username,
			email,
			password,
			"Bearer",
			accessToken,
			refreshToken,
			expiresIn,
			USER,
		)
		CheckError(err)
	}()

	response := &auth.SignupResponse{
		UserId:       userId,
		TokenType:    "Bearer",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    (int64)(expiresIn),
		Scope:        USER.String(),
	}
	end := time.Now().UnixMilli()
	logger.Printf("%v ms", (end - start))

	return response, nil
}

func (Server) Signin(
	ctx context.Context,
	request *auth.SigninRequest,
) (*auth.SigninResponse, error) {
	email := request.GetEmail()
	password := request.GetPassword()

	if email == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "One of the fields is empty")
	}

	userId, err := Db().GetUserId(ctx, email)

	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	accessToken := GenerateToken(userId)
	refreshToken := GenerateToken(userId)
	expiresIn := 60 * 60 * 24 * 60 * time.Second

	response := &auth.SigninResponse{
		UserId:       userId,
		TokenType:    "Bearer",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    (int64)(expiresIn),
		Scope:        USER.String(),
	}

	return response, nil
}
