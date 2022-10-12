package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app                           = kingpin.New("api-gateway", "restful json http api for the noted backend").DefaultEnvars()
	port                          = app.Flag("port", "http api port").Default("3000").Int16()
	environment                   = app.Flag("env", "production or development").Default(envIsProd).Enum(envIsProd, envIsDev)
	accountsServiceAddress        = app.Flag("accounts-service-addr", "the grpc address of the accounts service").Default("accounts:3000").String()
	recommendationsServiceAddress = app.Flag("recommendations-service-addr", "the grpc address of the recommendations service").Default("recommendations:3000").String()
)

const (
	envIsProd = "production"
	envIsDev  = "development"
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	s := server{}
	s.Init()

	s.Engine.Use(gin.Recovery())
	s.Engine.Use(s.LoggerMiddleware)
	s.Engine.Use(s.AccessControlMiddleware)
	s.Engine.Use(s.PreflightMiddleware)

	s.Engine.POST("/authenticate", s.accountsHandler.Authenticate)
	s.Engine.GET("/accounts", s.accountsHandler.List)
	s.Engine.GET("/accounts/:id", s.accountsHandler.Get)
	s.Engine.POST("/accounts", s.accountsHandler.Create)
	s.Engine.PATCH("/accounts/:id", s.accountsHandler.Update)
	s.Engine.DELETE("/accounts/:id", s.accountsHandler.Delete)

	s.Engine.POST("/groups", s.groupsHandler.Create)
	s.Engine.POST("/groups/:id/join", s.groupsHandler.Join)
	s.Engine.DELETE("/groups/:id", s.groupsHandler.Delete)
	s.Engine.PATCH("/groups/:id", s.groupsHandler.Update)
	s.Engine.GET("/groups/:id/members", s.groupsHandler.ListMembers)

	s.Engine.POST("/recommendations/keywords", s.recommendationsHandler.ExtractKeywords)

	s.Run()
	defer s.Close()
}
