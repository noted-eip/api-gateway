package main

import (
	pb "api-gateway/grpc/accountspb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(account pb.AccountsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		id := "none"
		email := "maxime@noted.io"

		req := &pb.Account{Id: id, Name: name, Email: email}
		res, err := account.CreateAccount(c, req)

		// Contact the server and print out its response.
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res),
		})
	}
}
