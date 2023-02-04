package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	notesv1 "api-gateway/protorepo/noted/notes/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app                    = kingpin.New("api-gateway", "restful json http api for the noted backend").DefaultEnvars()
	port                   = app.Flag("port", "http api port").Default("3000").Int16()
	environment            = app.Flag("env", "production or development").Default(envIsProd).Enum(envIsProd, envIsDev)
	accountsServiceAddress = app.Flag("accounts-service-addr", "the grpc address of the accounts service").Default("accounts:3000").String()
	notesServiceAddress    = app.Flag("notes-service-addr", "the grpc address of the notes service").Default("notes:3000").String()
)

const (
	envIsProd = "production"
	envIsDev  = "development"
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	logger := initLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	must(accountsv1.RegisterAccountsAPIHandlerFromEndpoint(ctx, mux, *accountsServiceAddress, opts))
	must(notesv1.RegisterGroupsAPIHandlerFromEndpoint(ctx, mux, *notesServiceAddress, opts))

	logger.Info("starting api-gateway", zap.Int16("port", *port))
	must(http.ListenAndServe(fmt.Sprint(":", *port), mux))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func initLogger() *zap.Logger {
	if *environment == envIsDev {
		logger, err := zap.NewDevelopment()
		must(err)
		return logger
	}
	logger, err := zap.NewProduction()
	must(err)

	return logger
}
