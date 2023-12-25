package handler

import (
	"lgtmeme_backend/api/_pkg/dto"
	"lgtmeme_backend/api/_pkg/service"
	"lgtmeme_backend/api/_pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostImageHandler(ctx *gin.Context) {
	var reqBody dto.PostImageReqBody

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errInfos := util.ValidateReqBody(err)
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	imageID := uuid.New().String()
	base64Image := reqBody.Image

	imageURL, err := service.UploadImageToStorage(ctx, imageID, base64Image)
	if err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}

	if err := service.InsertImage(ctx, imageID, imageURL, reqBody.Keyword); err != nil {
		if err := service.DeleteImageFromStorage(imageURL); err != nil {
			util.InternalServerErrorResponse(ctx, err)
			return
		}
		util.InternalServerErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"image_url": imageURL,
	})
}
