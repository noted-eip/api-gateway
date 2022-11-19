package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type invitesHandler struct {
	invitesClient accountsv1.InvitesAPIClient
}

func (h *invitesHandler) SendInvite(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.SendInviteRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.invitesClient.SendInvite(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *invitesHandler) GetInvite(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetInviteRequest{
		InviteId: c.Param("invite_id"),
	}

	res, err := h.invitesClient.GetInvite(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *invitesHandler) AcceptInvite(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.AcceptInviteRequest{
		InviteId: c.Param("invite_id"),
	}

	res, err := h.invitesClient.AcceptInvite(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *invitesHandler) DenyInvite(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.DenyInviteRequest{
		InviteId: c.Param("invite_id"),
	}

	res, err := h.invitesClient.DenyInvite(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *invitesHandler) ListInvites(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListInvitesRequest{
		SenderAccountId:    c.Query("sender_account_id"),
		RecipientAccountId: c.Query("recipient_account_id"),
		GroupId:            c.Query("group_id"),
		Limit:              queryAsInt32OrDefault(c, "limit", 0),
		Offset:             queryAsInt32OrDefault(c, "limit", 0),
	}

	res, err := h.invitesClient.ListInvites(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
