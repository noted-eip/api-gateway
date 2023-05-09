package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	notesv1 "api-gateway/protorepo/noted/notes/v1"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	srv := newServer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC APIs here.
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	must(accountsv1.RegisterAccountsAPIHandlerFromEndpoint(ctx, srv.mux, *accountsServiceAddress, opts))
	must(notesv1.RegisterGroupsAPIHandlerFromEndpoint(ctx, srv.mux, *notesServiceAddress, opts))
	must(notesv1.RegisterNotesAPIHandlerFromEndpoint(ctx, srv.mux, *notesServiceAddress, opts))
	must(notesv1.RegisterRecommendationsAPIHandlerFromEndpoint(ctx, srv.mux, *notesServiceAddress, opts))

	// Register routes
	srv.Engine.GET("/groups/:group_id/notes/:note_id/export", srv.notesHandler.ExportNote)

	srv.run()
}

type server struct {
	logger *zap.Logger
	mux    *runtime.ServeMux

	notesConn    *grpc.ClientConn
	notesClient  notesv1.NotesAPIClient
	notesHandler *notesHandler
	Engine       *gin.Engine
}

func newServer() *server {
	srv := &server{}
	srv.initNoteClient()
	srv.initLogger()
	srv.mux = runtime.NewServeMux(
		runtime.WithErrorHandler(srv.errorHandler),
		runtime.WithRoutingErrorHandler(srv.routingErrorHandler),
	)

	srv.Engine = gin.New()

	return srv
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (srv *server) initLogger() {
	var err error
	if *environment == envIsDev {
		srv.logger, err = zap.NewDevelopment()
		must(err)
	} else {
		srv.logger, err = zap.NewProduction()
		must(err)
	}
}

func (srv *server) initClientConn(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}

func (srv *server) initNoteClient() {
	srv.notesConn = srv.initClientConn(*notesServiceAddress)
	srv.notesClient = notesv1.NewNotesAPIClient(srv.notesConn)
	srv.notesHandler = &notesHandler{
		notesClient: srv.notesClient,
	}
}

type httpError struct {
	Error string `json:"error"`
}

func (srv *server) errorHandler(ctx context.Context, sm *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	st, ok := status.FromError(err)
	if !ok {
		srv.logger.Error("service replied with non-status error", zap.String("path", r.URL.Path), zap.String("method", r.Method))
		st = status.New(codes.Internal, "internal server error")
	}

	w.Header().Set("Content-Type", "application/json")
	httpStatus := runtime.HTTPStatusFromCode(st.Code())
	w.WriteHeader(httpStatus)

	bytes, err := json.Marshal(httpError{Error: st.Message()})
	if err != nil {
		srv.logger.Error("failed to marshal error response", zap.Any("status", st), zap.Error(err))
		return
	}

	if _, err := w.Write(bytes); err != nil {
		srv.logger.Error("failed to write error response", zap.Any("status", st), zap.Error(err))
	}
}

func (srv *server) routingErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	sterr := status.Error(codes.Internal, "unexpected routing error")
	switch httpStatus {
	case http.StatusBadRequest:
		sterr = status.Error(codes.InvalidArgument, "bad request")
	case http.StatusMethodNotAllowed:
		sterr = status.Error(codes.Unimplemented, "method not allowed")
	case http.StatusNotFound:
		sterr = status.Error(codes.NotFound, "endpoint does not exist")
	}
	srv.errorHandler(ctx, mux, marshaler, w, r, sterr)
}

func (srv *server) handleCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (srv *server) run() {
	srv.logger.Info("starting api-gateway", zap.Int16("port", *port))

	must(http.ListenAndServe(fmt.Sprint(":", *port), srv.handleCors(srv.mux)))
}
