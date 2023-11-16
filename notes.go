package main

import (
	notesv1 "api-gateway/protorepo/noted/notes/v1"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type notesHandler struct {
	notesClient notesv1.NotesAPIClient
	logger      *zap.Logger
}

func (h *notesHandler) CreateNote(w http.ResponseWriter, r *http.Request, pathParams map[string]string) /*(code int, contentType string, data []byte)*/ {
	bearer, err := authenticate(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err)
		return
	}

	// Read body and interpret it in a Create Note request struct
	var body notesv1.CreateNoteRequest
	if err := convertJsonToProto(r.Body, &body); err != nil {
		err := errors.New("wrong body: " + err.Error())
		writeError(w, http.StatusBadRequest, err)
		return
	}

	// Fetch group ID to complete Create Note request
	groupId, ok := pathParams["group_id"]
	if !ok {
		err := errors.New("group id error: " + err.Error())
		writeError(w, http.StatusBadRequest, err)
		return
	}
	body.GroupId = groupId

	// Fetch language to complete Create Note request
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "fr"
	}
	body.Lang = lang

	// Call the correct endpoint
	res, err := h.notesClient.CreateNote(contextWithGrpcBearer(context.Background(), bearer), &body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// Marshal result to JSON and respond
	resBytes, err := json.Marshal(res)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resBytes)

	if err != nil {
		h.logger.Error("Failed to create note : ", zap.Error(err))
	} else {
		h.logger.Info("Created note successfully")
	}
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
