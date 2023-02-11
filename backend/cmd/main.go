package main

import (
	"context"
	"fmt"
	"jangle/backend/auth"
	"jangle/backend/pkg"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
)

var logger *log.Logger = log.New(os.Stdout, "", log.LstdFlags)

func main() {
	cwd, _ := os.Getwd()
	err := pkg.LoadDotenv(
		path.Join(
			path.Dir(cwd),
			".env",
		),
	)
	pkg.CheckError(err)
	logger.Println("Loaded the environment variables")

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	postgresErr, mongoErr := pkg.Db().InitializeDB(ctx)
	pkg.CheckError(postgresErr)
	pkg.CheckError(mongoErr)
	logger.Println("Connected to the databases")

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	pkg.CheckError(err)
	logger.Printf("Listening on %s\n", lis.Addr().String())

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
	logger.Println("Initialised and registered gRPC servers")

	defer func() {
		grpcServer.GracefulStop()
		lis.Close()
		sentry.Flush(2 * time.Second)
	}()

	logger.Printf("Now serving on %s", lis.Addr().String())
	grpcServer.Serve(lis)
}
