package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type groupsHandler struct {
	groupsClient accountsv1.GroupsAPIClient
}

func (h *groupsHandler) Create(c *gin.Context) {
	body := &accountsv1.CreateGroupRequest{}
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
	body := &accountsv1.DeleteGroupRequest{
		Id: c.Param("id"),
	}

	res, err := h.groupsClient.DeleteGroup(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
