package main

import (
	recommendationsv1 "api-gateway/protorepo/noted/recommendations/v1"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type recommendationsHandler struct {
	recommendationsClient recommendationsv1.RecommendationsAPIClient
}

func (h *recommendationsHandler) Get(c *gin.Context) {
	body := &recommendationsv1.ExtractKeywordsRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}

	res, err := h.recommendationsClient.ExtractKeywords(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusOK, httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
