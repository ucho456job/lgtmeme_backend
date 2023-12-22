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
		"error_message": "validation failed for one or more fields",
		"errors":        errInfos,
	})
}

func InternalServerErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error_code":    "internal_server_error",
		"error_message": "internal server error occurs",
		"errors": []ErrInfo{{
			Field:   "",
			Tag:     "",
			Message: err.Error(),
		}},
	})
}
