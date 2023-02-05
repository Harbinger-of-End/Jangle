package main

import (
	"fmt"
	"jangle/backend/auth"
	"jangle/backend/pkg"
	"net"
	"os"
	"path"
	"time"

	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
)

func main() {
	cwd, _ := os.Getwd()
	err := pkg.LoadDotenv(
		path.Join(
			path.Dir(cwd),
			".env",
		),
	)
	pkg.CheckError(err)

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	pkg.CheckError(err)

	err = sentry.Init(
		sentry.ClientOptions{
			Dsn:              os.Getenv("SENTRY_DSN"),
			TracesSampleRate: 1.0,
		},
	)
	pkg.CheckError(err)

	grpcServer := grpc.NewServer()
	server := new(pkg.Server)
	auth.RegisterAuthenticationServer(grpcServer, server)

	defer func() {
		grpcServer.GracefulStop()
		lis.Close()
		defer sentry.Flush(2 * time.Second)
	}()

	grpcServer.Serve(lis)
}
