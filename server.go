package main

import (
	"api-gateway/grpc/accountspb"
	"api-gateway/grpc/groupspb"
	"api-gateway/grpc/recommendationspb"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	accountsConn *grpc.ClientConn

	accountsClient  accountspb.AccountsServiceClient
	accountsHandler *accountsHandler

	groupsClient  groupspb.GroupServiceClient
	groupsHandler *groupsHandler

	recommendationsConn    *grpc.ClientConn
	recommendationsClient  recommendationspb.RecommendationsServiceClient
	recommendationsHandler *recommendationsHandler

	logger  *zap.Logger
	slogger *zap.SugaredLogger

	Engine *gin.Engine
}

func (s *server) Init() {
	s.accountsConn = s.initClientConn(*accountsServiceAddress)
	s.accountsClient = accountspb.NewAccountsServiceClient(s.accountsConn)
	s.accountsHandler = &accountsHandler{
		accountsClient: s.accountsClient,
	}

	s.groupsClient = groupspb.NewGroupServiceClient(s.accountsConn)
	s.groupsHandler = &groupsHandler{
		groupsClient: s.groupsClient,
	}

	s.recommendationsConn = s.initClientConn(*recommendationsServiceAddress)
	s.recommendationsClient = recommendationspb.NewRecommendationsServiceClient(s.accountsConn)
	s.recommendationsHandler = &recommendationsHandler{
		recommendationsClient: s.recommendationsClient,
	}

	s.initLogger()

	gin.SetMode(gin.ReleaseMode)
	s.Engine = gin.New()
}

func (s *server) Run() {
	s.slogger.Infof("service running on :%d", *port)
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
		s.logger, err = zap.NewDevelopment(zap.WithCaller(false))
		if err != nil {
			panic(err)
		}
	} else {
		s.logger, err = zap.NewProduction(zap.WithCaller(false))
		if err != nil {
			panic(err)
		}
	}
	s.slogger = s.logger.Sugar()
}
