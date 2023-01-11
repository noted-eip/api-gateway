package main

import (
	accountsv1 "api-gateway/protorepo/noted/accounts/v1"
	notesv1 "api-gateway/protorepo/noted/notes/v1"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type notesHandler struct {
	notesClient  notesv1.NotesAPIClient
	groupsClient accountsv1.GroupsAPIClient
}

func (h *notesHandler) CreateNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	createNoteBody := &notesv1.CreateNoteRequest{}
	if err := readRequestBody(c, createNoteBody); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	note, err := h.notesClient.CreateNote(contextWithGrpcBearer(context.Background(), bearer), createNoteBody)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	addGroupNoteBody := &accountsv1.AddGroupNoteRequest{
		GroupId: c.Param("group_id"),
		NoteId:  note.Note.Id,
		Title:   note.Note.Title,
	}

	_, err = h.groupsClient.AddGroupNote(contextWithGrpcBearer(context.Background(), bearer), addGroupNoteBody)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, note)
}

func (h *notesHandler) GetNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.GetNoteRequest{
		Id: c.Param("note_id"),
	}

	res, err := h.notesClient.GetNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *notesHandler) UpdateNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.UpdateNoteRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.Id = c.Param("note_id")

	res, err := h.notesClient.UpdateNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *notesHandler) DeleteNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.DeleteNoteRequest{
		Id: c.Param("note_id"),
	}

	res, err := h.notesClient.DeleteNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *notesHandler) ListNotes(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.ListNotesRequest{
		AuthorId: c.Query("author_id"),
	}

	res, err := h.notesClient.ListNotes(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *notesHandler) ExportNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	formatMap := map[string]struct {
		Format      notesv1.NoteExportFormat
		ContentType string
	}{
		"":    {Format: notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_INVALID, ContentType: ""},
		"md":  {Format: notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_MARKDOWN, ContentType: "text/markdown; charset=UTF-8"},
		"pdf": {Format: notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_PDF, ContentType: "application/pdf"},
	}

	format, ok := formatMap[c.Query("format")]
	if !ok {
		err := errors.New("unknow export format")
		writeError(c, http.StatusBadRequest, err)
		return
	}

	body := &notesv1.ExportNoteRequest{
		ExportFormat: notesv1.NoteExportFormat(format.Format),
	}
	body.NoteId = c.Param("note_id")

	res, err := h.notesClient.ExportNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, format.ContentType, res.File)
}

func (h *notesHandler) InsertBlock(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.InsertBlockRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.NoteId = c.Param("note_id")

	res, err := h.notesClient.InsertBlock(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *notesHandler) UpdateBlock(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.UpdateBlockRequest{}
	if err := readRequestBody(c, body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.Id = c.Param("block_id")

	res, err := h.notesClient.UpdateBlock(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}

func (h *notesHandler) DeleteBlock(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.DeleteBlockRequest{
		Id: c.Param("block_id"),
	}

	res, err := h.notesClient.DeleteBlock(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	writeResponse(c, res)
}
