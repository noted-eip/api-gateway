package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
)

type accountsHandler struct {
	accountsClient accountsv1.AccountsAPIClient
}

func (h *accountsHandler) Authenticate(c *gin.Context) {
	body := &accountsv1.AuthenticateRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.accountsClient.Authenticate(context.Background(), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) Create(c *gin.Context) {
	body := &accountsv1.CreateAccountRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.accountsClient.CreateAccount(context.Background(), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) Get(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetAccountRequest{
		Id: c.Param("id"),
	}

	res, err := h.accountsClient.GetAccount(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "unimplemented"})
}

func (h *accountsHandler) Update(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateAccountRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}
	body.Account.Id = c.Param("id")

	res, err := h.accountsClient.UpdateAccount(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) Delete(c *gin.Context) {
	bearer, err := h.authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.DeleteAccountRequest{
		Id: c.Param("id"),
	}

	res, err := h.accountsClient.DeleteAccount(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// authenticate fetches the bearer string from the authorization header or
// returns an error if it is missing.
func (h *accountsHandler) authenticate(c *gin.Context) (string, error) {
	bearer := c.GetHeader(httpAuthorizationHeader)
	if bearer == "" {
		return "", ErrUnauthenticated
	}
	return bearer, nil
}
