package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrInfo struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func ValidationFailedResponse(ctx *gin.Context, errInfos []ErrInfo) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error_code":    "validation_failed",
		"error_message": "Validation failed for one or more fields.",
		"errors":        errInfos,
	})
}
