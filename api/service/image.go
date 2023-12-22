package service

import (
	"lgtmeme_backend/api/config"
	"lgtmeme_backend/api/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func InsertImage(ctx *gin.Context, ID string, url string, keyword string) error {
	newImage := model.Image{
		ID:        ID,
		URL:       url,
		Keyword:   keyword,
		CreatedAt: time.Now(),
	}

	result := config.DB.Table("images").Create(&newImage)
	if result.Error != nil {
		log.Printf("Failed to insert image: %v", result.Error)
		return result.Error
	}

	return nil
}
