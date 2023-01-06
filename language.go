package main

import (
	languagev1 "api-gateway/protorepo/noted/language/v1"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type languageHandler struct {
	languageClient languagev1.LanguageAPIClient
}

func (h *languageHandler) ExtractKeywords(c *gin.Context) {
	body := &languagev1.ExtractKeywordsRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		writeError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.languageClient.ExtractKeywords(context.Background(), body)
	if err != nil {
		writeError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
