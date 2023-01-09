package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type conversationsHandler struct {
	conversationsClient accountsv1.ConversationsAPIClient
}

func (h *conversationsHandler) CreateConversation(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.CreateConversationRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.GroupId = c.Query("group_id")

	res, err := h.conversationsClient.CreateConversation(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) GetConversation(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetConversationRequest{
		ConversationId: c.Param("conversation_id"),
	}

	res, err := h.conversationsClient.GetConversation(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) ListConversations(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListConversationsRequest{
		GroupId: c.Query("group_id"),
	}

	res, err := h.conversationsClient.ListConversations(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) UpdateConversation(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateConversationRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.ConversationId = c.Param("conversation_id")

	res, err := h.conversationsClient.UpdateConversation(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) DeleteConversation(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.DeleteConversationRequest{
		ConversationId: c.Param("conversation_id"),
	}

	res, err := h.conversationsClient.DeleteConversation(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) SendConversationMessage(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.SendConversationMessageRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	body.ConversationId = c.Param("conversation_id")

	res, err := h.conversationsClient.SendConversationMessage(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) DeleteConversationMessage(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.DeleteConversationMessageRequest{
		MessageId:      c.Param("message_id"),
		ConversationId: c.Param("conversation_id"),
	}

	res, err := h.conversationsClient.DeleteConversationMessage(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) ListConversationMessages(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListConversationMessagesRequest{
		ConversationId: c.Param("conversation_id"),
		Limit:          queryAsInt32OrDefault(c, "limit", 0),
		Offset:         queryAsInt32OrDefault(c, "offset", 0),
	}

	res, err := h.conversationsClient.ListConversationMessages(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) UpdateConversationMessage(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateConversationMessageRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.MessageId = c.Param("message_id")
	body.ConversationId = c.Param("conversation_id")

	res, err := h.conversationsClient.UpdateConversationMessage(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *conversationsHandler) GetConversationMessage(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetConversationMessageRequest{
		ConversationId: c.Param("conversation_id"),
		MessageId:      c.Param("message_id"),
	}

	res, err := h.conversationsClient.GetConversationMessage(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}
