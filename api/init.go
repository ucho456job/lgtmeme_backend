package api

import (
	"lgtmeme_backend/api/_pkg/config"
	"lgtmeme_backend/api/_pkg/handler"
	"lgtmeme_backend/api/_pkg/middleware"
	"lgtmeme_backend/api/_pkg/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func initGin() *gin.RouterGroup {
	engine = gin.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("base64image", util.IsValidBase64Image)
		v.RegisterValidation("imagesize", util.IsValidImageSize)
	}
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
