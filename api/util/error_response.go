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

// -------------------
// 400
// -------------------
func ValidationFailedResponse(ctx *gin.Context, errInfos []ErrInfo) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error_code":    "validation_failed",
		"error_message": "Validation failed for one or more fields.",
		"errors":        errInfos,
	})
}

// -------------------
// 500
// -------------------
func InternalServerErrorResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error_code":    "internal_server_error",
		"error_message": "",
		"erros":         ErrInfo{},
	})
}
