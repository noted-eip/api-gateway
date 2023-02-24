package main

import (
	notesv1 "api-gateway/protorepo/noted/notes/v1"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type notesHandler struct {
	notesClient notesv1.NotesAPIClient
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
