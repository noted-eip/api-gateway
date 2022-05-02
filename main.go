package main

import (
	"log"

	pb "api-gateway/grpc/accountspb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	AccountClient := pb.NewAccountsServiceClient(conn)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up a http server.
	r := gin.Default()

	r.GET("/rest/n/:name", account(AccountClient))

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
