package main

import (
	notesv1 "api-gateway/protorepo/noted/notes/v1"
	"context"
	"net/http"

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
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
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
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
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
		AuthorId: c.Param("author_id"),
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

	body := &notesv1.ExportNoteRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}
	body.NoteId = c.Param("note_id")

	res, err := h.notesClient.ExportNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *notesHandler) InsertBlock(c *gin.Context) {
	bearer, err := authenticate(c)
	if err != nil {
		writeError(c, http.StatusUnauthorized, err)
		return
	}

	body := &notesv1.InsertBlockRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
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
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
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
