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
	notesServiceAddress           = app.Flag("notes-service-addr", "the grpc address of the notes service").Default("notes:3000").String()
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

	// Accounts
	s.Engine.POST("/accounts", s.accountsHandler.CreateAccount)
	s.Engine.GET("/accounts/:account_id", s.accountsHandler.GetAccount)
	s.Engine.PATCH("/accounts/:account_id", s.accountsHandler.UpdateAccount)
	s.Engine.DELETE("/accounts/:account_id", s.accountsHandler.DeleteAccount)
	s.Engine.GET("/accounts", s.accountsHandler.ListAccounts)
	s.Engine.POST("/authenticate", s.accountsHandler.Authenticate)

	// Groups
	s.Engine.POST("/groups", s.groupsHandler.CreateGroup)
	s.Engine.GET("/groups/:group_id", s.groupsHandler.GetGroup)
	s.Engine.PATCH("/groups/:group_id", s.groupsHandler.UpdateGroup)
	s.Engine.DELETE("/groups/:group_id", s.groupsHandler.DeleteGroup)
	s.Engine.GET("/groups", s.groupsHandler.ListGroups)

	// Group Members
	s.Engine.GET("/groups/:group_id/members/:member_id", s.groupsHandler.GetGroupMember)
	s.Engine.PATCH("/groups/:group_id/members/:member_id", s.groupsHandler.UpdateGroupMember)
	s.Engine.DELETE("/groups/:group_id/members/:member_id", s.groupsHandler.RemoveGroupMember)
	s.Engine.GET("/groups/:group_id/members", s.groupsHandler.ListGroupMembers)

	// Group Notes
	s.Engine.POST("/groups/:group_id/notes", s.groupsHandler.AddGroupNote)
	s.Engine.GET("/groups/:group_id/notes/:note_id", s.groupsHandler.GetGroupNote)
	s.Engine.PATCH("/groups/:group_id/notes/:note_id", s.groupsHandler.UpdateGroupNote)
	s.Engine.DELETE("/groups/:group_id/notes/:note_id", s.groupsHandler.RemoveGroupNote)
	s.Engine.GET("/groups/:group_id/notes", s.groupsHandler.ListGroupNotes)

	// Invites
	s.Engine.POST("/invites", s.invitesHandler.SendInvite)
	s.Engine.GET("/invites/:invite_id", s.invitesHandler.GetInvite)
	s.Engine.POST("/invites/:invite_id/accept", s.invitesHandler.AcceptInvite)
	s.Engine.POST("/invites/:invite_id/deny", s.invitesHandler.DenyInvite)
	s.Engine.GET("/invites", s.invitesHandler.ListInvites)

	// Notes
	s.Engine.GET("/notes/:note_id", s.notesHandler.GetNote)
	s.Engine.POST("/notes", s.notesHandler.CreateNote)
	s.Engine.PATCH("/notes/:note_id", s.notesHandler.UpdateNote)
	s.Engine.DELETE("/notes/:note_id", s.notesHandler.DeleteNote)
	s.Engine.GET("/notes/?author_id=", s.notesHandler.ListNotes)
	s.Engine.GET("/notes/:note_id/export/?format=", s.notesHandler.ExportNote)

	// Blocks
	s.Engine.POST("/notes/:note_id/blocks", s.notesHandler.InsertBlock)
	s.Engine.PATCH("/notes/:note_id/blocks/:block_id", s.notesHandler.UpdateBlock)
	s.Engine.DELETE("/notes/:note_id/blocks/:block_id", s.notesHandler.DeleteBlock)

	// Recommendations
	s.Engine.POST("/recommendations/keywords", s.recommendationsHandler.ExtractKeywords)

	s.Run()
	defer s.Close()
}
