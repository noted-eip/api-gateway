package main

import (
	notesv1 "api-gateway/protorepo/noted/notes/v1"
	"bytes"
	"context"
	"errors"
	"net/http"
	"time"
)

type notesHandler struct {
	notesClient notesv1.NotesAPIClient
}

func (h *notesHandler) ExportNote(w http.ResponseWriter, r *http.Request, pathParams map[string]string) /*(code int, contentType string, data []byte)*/ {
	bearer, err := authenticate(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err)
		return
	}

	formatMap := map[string]notesv1.NoteExportFormat{
		"":    notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_INVALID,
		"md":  notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_MARKDOWN,
		"pdf": notesv1.NoteExportFormat_NOTE_EXPORT_FORMAT_PDF,
	}

	fileType := r.URL.Query().Get("format")

	format, ok := formatMap[fileType]
	if !ok {
		err := errors.New("unknow export format")
		writeError(w, http.StatusBadRequest, err)
		return
	}

	body := &notesv1.ExportNoteRequest{
		NoteId:       pathParams["note_id"],
		GroupId:      pathParams["group_id"],
		ExportFormat: notesv1.NoteExportFormat(format),
	}

	res, err := h.notesClient.ExportNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	fileName := "note." + fileType
	w.Header().Set("Content-Disposition", "attachment;filename="+fileName) // Headers that tells the browser to download the served file with the name note.pdf
	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(res.File))
}
