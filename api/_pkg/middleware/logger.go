package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func WriteLog(ctx *gin.Context) {
	fmt.Println("middleware: WriteLog")
}
