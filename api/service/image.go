package service

import (
	"lgtmeme_backend/api/config"
	"lgtmeme_backend/api/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TABLE_NAME = "images"

func InsertImage(ctx *gin.Context, url string, keyword string) (imageID string, err error) {
	imageID = uuid.New().String()
	newImage := model.Image{
		ID:        imageID,
		URL:       url,
		Keyword:   keyword,
		CreatedAt: time.Now(),
	}

	result := config.DB.Table(TABLE_NAME).Create(&newImage)
	if result.Error != nil {
		log.Printf("Failed to insert image: %v", result.Error)
		return imageID, result.Error
	}

	return imageID, nil
}
