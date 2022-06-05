package main

import (
	"api-gateway/grpc/recommendationspb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type recommendationsHandler struct {
	recommendationsClient recommendationspb.RecommendationsServiceClient
}

func (h *recommendationsHandler) Get(c *gin.Context) {
	body := &recommendationspb.ExtractKeywordsRequest{}
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
