package service

import (
	"lgtmeme_backend/api/_pkg/config"
	"lgtmeme_backend/api/_pkg/entity"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func SelectImages(ctx *gin.Context, page int, keyword string, sort string, favoriteIDs []string, authCheck bool) (images []struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}, err error) {
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

func ExistsImage(ctx *gin.Context, ID string) (exists bool, err error) {
	var count int64
	result := config.DB.Table("images").Where("id = ?", ID).Count(&count)
	if result.Error != nil {
		log.Printf("Failed to check image existence: %v", result.Error)
		return false, result.Error
	}

	return count > 0, nil
}

func UpdateImage(ctx *gin.Context, ID string, requestType string) error {
	var updateData map[string]interface{}

	switch requestType {
	case "used":
		updateData = map[string]interface{}{"used_count": gorm.Expr("used_count + ?", 1)}
	case "reporting":
		updateData = map[string]interface{}{"reported": true}
	case "confirmed":
		updateData = map[string]interface{}{"confirmed": true}
	}

	result := config.DB.Table("images").Where("id = ?", ID).Updates(updateData)
	if result.Error != nil {
		log.Printf("Failed to update image: %v", result.Error)
		return result.Error
	}

	return nil
}
