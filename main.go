package main

import (
	"fmt"
	"log"
	"net/http"

	pb "api-gateway/grpc/accountspb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	AccountClient := pb.NewAccountsServiceClient(conn)

	// Set up a http server.
	r := gin.Default()
	r.GET("/rest/n/:name", func(c *gin.Context) {
		name := c.Param("name")
		id := "none"
		email := "maxime@noted.io"

		// Contact the server and print out its response.
		req := &pb.Account{Id: id, Name: name, Email: email}
		res, err := AccountClient.CreateAccount(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res),
		})

	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
