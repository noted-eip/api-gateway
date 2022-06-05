package main

import (
	"api-gateway/grpc/groupspb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type groupsHandler struct {
	groupsClient groupspb.GroupServiceClient
}

func (h *groupsHandler) Create(c *gin.Context) {
	body := &groupspb.CreateGroupRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}

	res, err := h.groupsClient.CreateGroup(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *groupsHandler) Delete(c *gin.Context) {
	body := &groupspb.GroupFilterRequest{
		Id: c.Param("id"),
	}

	res, err := h.groupsClient.DeleteGroup(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
