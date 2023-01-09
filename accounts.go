package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountsHandler struct {
	accountsClient accountsv1.AccountsAPIClient
}

func (h *accountsHandler) CreateAccount(c *gin.Context) {
	body := &accountsv1.CreateAccountRequest{}
	if err := readRequestBody(c, body); err != nil {
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

func (h *accountsHandler) GetAccount(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetAccountRequest{
		Id:    c.Param("account_id"),
		Email: c.Param("email"),
	}

	res, err := h.accountsClient.GetAccount(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) ListAccounts(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListAccountsRequest{
		Limit:  queryAsInt32OrDefault(c, "limit", 0),
		Offset: queryAsInt32OrDefault(c, "offset", 0),
	}

	res, err := h.accountsClient.ListAccounts(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) UpdateAccount(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateAccountRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.Account.Id = c.Param("account_id")

	res, err := h.accountsClient.UpdateAccount(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) DeleteAccount(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.DeleteAccountRequest{
		Id: c.Param("account_id"),
	}

	res, err := h.accountsClient.DeleteAccount(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *accountsHandler) Authenticate(c *gin.Context) {
	body := &accountsv1.AuthenticateRequest{}
	if err := readRequestBody(c, body); err != nil {
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
