package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	recommendationsv1 "api-gateway/protorepo/noted/recommendations/v1"
	"net/http"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	accountsConn *grpc.ClientConn

	accountsClient  accountsv1.AccountsAPIClient
	accountsHandler *accountsHandler

	groupsClient  accountsv1.GroupsAPIClient
	groupsHandler *groupsHandler

	recommendationsConn    *grpc.ClientConn
	recommendationsClient  recommendationsv1.RecommendationsAPIClient
	recommendationsHandler *recommendationsHandler

	logger  *zap.Logger
	slogger *zap.SugaredLogger

	Engine *gin.Engine
}

func (s *server) Init() {
	s.accountsConn = s.initClientConn(*accountsServiceAddress)
	s.accountsClient = accountsv1.NewAccountsAPIClient(s.accountsConn)
	s.accountsHandler = &accountsHandler{
		accountsClient: s.accountsClient,
	}

	s.groupsClient = accountsv1.NewGroupsAPIClient(s.accountsConn)
	s.groupsHandler = &groupsHandler{
		groupsClient: s.groupsClient,
	}

	s.recommendationsConn = s.initClientConn(*recommendationsServiceAddress)
	s.recommendationsClient = recommendationsv1.NewRecommendationsAPIClient(s.recommendationsConn)
	s.recommendationsHandler = &recommendationsHandler{
		recommendationsClient: s.recommendationsClient,
	}

	s.initLogger()

	gin.SetMode(gin.ReleaseMode)
	s.Engine = gin.New()
}

func (s *server) Run() {
	s.slogger.Infof("api-gateway running on :%d", *port)
	err := s.Engine.Run(fmt.Sprint(":", *port))
	if err != nil {
		panic(err)
	}
}

func (s *server) LoggerMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()
	if c.Writer.Status() > 499 {
		s.logger.Error("failed http request",
			zap.Int("code", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("endpoint", c.Request.URL.Path),
			zap.Duration("duration", time.Since(start)),
		)
		return
	} else if c.Writer.Status() > 399 {
		s.logger.Warn("invalid http request",
			zap.Int("code", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("endpoint", c.Request.URL.Path),
			zap.Duration("duration", time.Since(start)),
		)
		return
	}
	s.logger.Info("http request",
		zap.Int("code", c.Writer.Status()),
		zap.String("method", c.Request.Method),
		zap.String("endpoint", c.Request.URL.Path),
		zap.Duration("duration", time.Since(start)),
	)
}

func (s *server) AccessControlMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Next()
}

// TODO: Only invoke this middleware in registered routes.
func (s *server) PreflightMiddleware(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Status(http.StatusOK)
	}
	c.Next()
}

func (s *server) Close() {
	s.logger.Info("graceful shutdown")
	s.accountsConn.Close()
	s.logger.Sync()
}

func (s *server) initClientConn(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}

func (s *server) initLogger() {
	var err error
	if *environment == envIsDev {
		s.logger, err = zap.NewDevelopment(zap.WithCaller(false), zap.AddStacktrace(zap.PanicLevel))
		if err != nil {
			panic(err)
		}
	} else {
		s.logger, err = zap.NewProduction(zap.WithCaller(false), zap.AddStacktrace(zap.PanicLevel))
		if err != nil {
			panic(err)
		}
	}
	s.slogger = s.logger.Sugar()
}
