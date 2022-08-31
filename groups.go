package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type groupsHandler struct {
	groupsClient accountsv1.GroupsAPIClient
}

func (h *groupsHandler) Create(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.CreateGroupRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.groupsClient.CreateGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *groupsHandler) Delete(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.DeleteGroupRequest{
		Id: c.Param("id"),
	}

	res, err := h.groupsClient.DeleteGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *groupsHandler) ListMembers(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListGroupMembersRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}

	body.Id = c.Param("id")

	res, err := h.groupsClient.ListGroupMembers(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *groupsHandler) Update(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateGroupRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}
	body.Group.Id = c.Param("id")
	res, err := h.groupsClient.UpdateGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *groupsHandler) Join(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.JoinGroupRequest{}
	body.Id = c.Param("id")

	res, err := h.groupsClient.JoinGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// authenticate fetches the bearer string from the authorization header or
// returns an error if it is missing.
func (h *groupsHandler) authenticate(c *gin.Context) (string, error) {
	bearer := c.GetHeader(httpAuthorizationHeader)
	if bearer == "" {
		return "", ErrUnauthenticated
	}
	return bearer, nil
}
