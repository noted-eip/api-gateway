package main

import (
	notesv1 "api-gateway/protorepo/noted/notes/v1"
	"context"
	"encoding/json"
	"errors"
	"net/http"
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

	format, ok := formatMap[r.URL.Query().Get("format")]
	if !ok {
		err := errors.New("unknow export format")
		writeError(w, http.StatusBadRequest, err)
		return
	}

	body := &notesv1.ExportNoteRequest{
		ExportFormat: notesv1.NoteExportFormat(format),
	}

	body.NoteId = r.FormValue("note_id")

	res, err := h.notesClient.ExportNote(contextWithGrpcBearer(context.Background(), bearer), body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	resp := make(map[string]string)
	resp["File"] = string(res.File)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
