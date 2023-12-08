package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func myRoute(rg *gin.RouterGroup) {
	rg.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Health check is ok.")
	})
}

func init() {
	engine = gin.New()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", []byte("<h1>Hello World</h1>"))
	})
	rg := engine.Group("/api")
	myRoute(rg)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	engine.ServeHTTP(w, r)
}
