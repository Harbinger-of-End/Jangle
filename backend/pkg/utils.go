package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Scope int64

const (
	USER Scope = 1 << iota
	BOT
	IDENTITY_READ
	GUILDS_JOIN
	GUILDS_READ
	GUILDS_MEMBER_READ
)

var generatedIds int64 = 0

const EPOCH int64 = 1672531200000

func (s Scope) String() string {
	switch s {
	case USER:
		return "user"
	case BOT:
		return "bot"
	case IDENTITY_READ:
		return "identity_read"
	case GUILDS_JOIN:
		return "guilds_join"
	case GUILDS_READ:
		return "guilds_read"
	case GUILDS_MEMBER_READ:
		return "guilds_member_read"
	default:
		return ""
	}
}

func GenerateSnowflake() string {
	workerId, err := strconv.Atoi(os.Getenv("WORKER_ID"))
	CheckError(err)
	processId, err := strconv.Atoi(os.Getenv("PROCESS_ID"))
	CheckError(err)

	snowflakeBase := int64((workerId << 17) + (processId << 12))
	atomic.AddInt64(
		&generatedIds,
		(1+generatedIds)%4096,
	)
	snowflake := snowflakeBase + (time.Now().UnixMilli()-EPOCH)<<22 + generatedIds

	return fmt.Sprint(snowflake)
}

func GenerateJWT(claims map[string]any) (string, error) {
	newClaims := (jwt.MapClaims)(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseJWT(tokenString string) map[string]any {
	token, err := jwt.Parse(
		tokenString,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	CheckError(err)

	return token.Claims.(jwt.MapClaims)
}

func GenerateToken(userId string) string {
	userIdBytes := []byte(userId)
	userIdPart := make([]byte, base64.StdEncoding.EncodedLen(len(userIdBytes)))
	base64.StdEncoding.Encode(userIdPart, userIdBytes)
	uniqueBytes := []byte(fmt.Sprint(time.Now().UnixMilli()))
	uniquePart := make([]byte, base64.StdEncoding.EncodedLen(len(uniqueBytes)))
	base64.StdEncoding.Encode(uniquePart, uniqueBytes)
	randomBytes := []byte(uuid.NewString())
	randomPart := make([]byte, base64.StdEncoding.EncodedLen(len(randomBytes)))
	base64.StdEncoding.Encode(randomPart, randomBytes)

	hash := sha256.New()
	hash.Write(uniquePart)
	hash.Write([]byte("."))
	hash.Write(randomPart)
	hashed := hash.Sum(nil)

	return fmt.Sprintf("%x.%x", userIdPart, hashed)
}
