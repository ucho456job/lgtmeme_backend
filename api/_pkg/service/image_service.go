package service

import (
	"lgtmeme_backend/api/_pkg/config"
	"lgtmeme_backend/api/_pkg/entity"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func InsertImage(ctx *gin.Context, ID string, url string, keyword string) error {
	newImage := entity.Image{
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

type selectImages []struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func SelectImages(ctx *gin.Context, page int, keyword string, sort string, favoriteIDs []string, authCheck bool) (images selectImages, err error) {
	q := config.DB
	q = q.Select("id", "url")
	q = q.Table("images")
	if len(favoriteIDs) > 0 {
		q = q.Where("id IN ?", favoriteIDs)
	}
	if keyword != "" {
		q = q.Where("keyword LIKE ?", "%"+keyword+"%")
	}
	if authCheck {
		q = q.Where("confirmed = ?", false)
		q = q.Where("reported = ?", true)
	}
	if sort == "popular" {
		q = q.Order("used_count DESC, created_at DESC")
	} else {
		q = q.Order("created_at DESC")
	}
	q = q.Offset(page * config.MAX_IMAGES_FETCH_COUNT).Limit(config.MAX_IMAGES_FETCH_COUNT)

	result := q.Find(&images)
	if result.Error != nil {
		log.Printf("Failed to select images: %v", result.Error)
		return images, result.Error
	}
	return images, nil
}
