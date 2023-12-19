package api

import (
	"lgtmeme_backend/api/config"
	"lgtmeme_backend/api/handler"
	"lgtmeme_backend/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	engine *gin.Engine
)

func init() {
	loadDotenv()
	config.InitDB()
	rg := initGin()
	initRoute(rg)
}

func loadDotenv() {
	err := godotenv.Load("env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func initGin() *gin.RouterGroup {
	engine = gin.New()
	engine.Use(middleware.WriteLog)
	rg := engine.Group("/api")
	return rg
}

func initRoute(rg *gin.RouterGroup) {
	rg.GET("/health", handler.Health)

	rg.POST("/images", handler.PostImageHandler)
	rg.GET("/images", handler.GetImageHandler)
	rg.PATCH("/images", handler.PatchImageHandler)
	rg.DELETE("/images", handler.DeleteImageHandler)
}
