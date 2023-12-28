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

func ResourceNotFoundResponse(ctx *gin.Context, errMsg string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error_code":    "resource_not_found",
		"error_message": "resource not found",
		"erros": []ErrInfo{{
			Field:   "",
			Tag:     "",
			Message: errMsg,
		}},
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
