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

func (h *groupsHandler) CreateGroup(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.CreateGroupRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.groupsClient.CreateGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) GetGroup(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetGroupRequest{
		GroupId: c.Param("group_id"),
	}

	res, err := h.groupsClient.GetGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) DeleteGroup(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.DeleteGroupRequest{
		GroupId: c.Param("group_id"),
	}

	res, err := h.groupsClient.DeleteGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) UpdateGroup(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateGroupRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	body.Group.Id = c.Param("group_id")
	res, err := h.groupsClient.UpdateGroup(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) ListGroups(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListGroupsRequest{
		AccountId: c.Query("account_id"),
		Limit:     queryAsInt32OrDefault(c, "limit", 0),
		Offset:    queryAsInt32OrDefault(c, "offset", 0),
	}

	res, err := h.groupsClient.ListGroups(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) GetGroupMember(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetGroupMemberRequest{
		GroupId:   c.Param("group_id"),
		AccountId: c.Param("member_id"),
	}

	res, err := h.groupsClient.GetGroupMember(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) UpdateGroupMember(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateGroupMemberRequest{
		Member: &accountsv1.GroupMember{},
	}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	body.GroupId = c.Param("group_id")
	body.Member.AccountId = c.Param("member_id")

	res, err := h.groupsClient.UpdateGroupMember(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) RemoveGroupMember(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.RemoveGroupMemberRequest{
		GroupId:   c.Param("group_id"),
		AccountId: c.Param("member_id"),
	}

	res, err := h.groupsClient.RemoveGroupMember(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) ListGroupMembers(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListGroupMembersRequest{
		GroupId: c.Param("group_id"),
		Limit:   queryAsInt32OrDefault(c, "limit", 0),
		Offset:  queryAsInt32OrDefault(c, "offset", 0),
	}

	res, err := h.groupsClient.ListGroupMembers(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) AddGroupNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.AddGroupNoteRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	body.GroupId = c.Param("group_id")

	res, err := h.groupsClient.AddGroupNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) GetGroupNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.GetGroupNoteRequest{
		GroupId: c.Param("group_id"),
		NoteId:  c.Param("note_id"),
	}

	res, err := h.groupsClient.GetGroupNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) UpdateGroupNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.UpdateGroupNoteRequest{
		Note: &accountsv1.GroupNote{},
	}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	body.GroupId = c.Param("group_id")
	body.Note.NoteId = c.Param("note_id")

	res, err := h.groupsClient.UpdateGroupNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) RemoveGroupNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.RemoveGroupNoteRequest{
		GroupId: c.Param("group_id"),
		NoteId:  c.Param("note_id"),
	}

	res, err := h.groupsClient.RemoveGroupNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *groupsHandler) ListGroupNotes(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &accountsv1.ListGroupNotesRequest{
		GroupId:         c.Param("group_id"),
		AuthorAccountId: c.Query("author_account_id"),
		FolderId:        c.Query("folder_id"),
		Limit:           queryAsInt32OrDefault(c, "limit", 0),
		Offset:          queryAsInt32OrDefault(c, "offset", 0),
	}

	res, err := h.groupsClient.ListGroupNotes(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}
