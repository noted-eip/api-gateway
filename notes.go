package main

import (
	notesv1 "api-gateway/protorepo/noted/notes/v1"
	"context"
	"errors"
	"io/ioutil"
	"net/http"

	protobuf "google.golang.org/protobuf/encoding/protojson"

	"github.com/gin-gonic/gin"
)

type notesHandler struct {
	notesClient notesv1.NotesAPIClient
}

func (h *notesHandler) CreateNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.CreateNoteRequest{}

	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	protobuf.Unmarshal(requestBody, body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	res, err := h.notesClient.CreateNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
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

	c.JSON(http.StatusOK, res)
}

func (h *notesHandler) UpdateNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.UpdateNoteRequest{}

	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	protobuf.Unmarshal(requestBody, body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}
	body.Id = c.Param("note_id")

	res, err := h.notesClient.UpdateNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
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

	c.JSON(http.StatusOK, res)
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

	c.JSON(http.StatusOK, res)
}

func (h *notesHandler) ExportNote(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	formatMap := map[string]notesv1.NoteExportFormat{
		"":    notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_INVALID,
		"md":  notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_MARKDOWN,
		"pdf": notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_PDF,
	}

	format, ok := formatMap[c.Query("format")]
	if !ok {
		err := errors.New("unknow export format")
		writeError(c, http.StatusBadRequest, err)
		return
	}

	body := &notesv1.ExportNoteRequest{
		ExportFormat: notesv1.NoteExportFormat(format),
	}
	body.NoteId = c.Param("note_id")

	res, err := h.notesClient.ExportNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, "File", res.File)
}

func (h *notesHandler) InsertBlock(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.InsertBlockRequest{}
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	protobuf.Unmarshal(requestBody, body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}
	body.NoteId = c.Param("note_id")

	res, err := h.notesClient.InsertBlock(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *notesHandler) UpdateBlock(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.UpdateBlockRequest{}
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	protobuf.Unmarshal(requestBody, body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}
	body.Id = c.Param("block_id")

	res, err := h.notesClient.UpdateBlock(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
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

	c.JSON(http.StatusOK, res)
}
