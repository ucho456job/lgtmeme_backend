package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Server is up and running.")
}
