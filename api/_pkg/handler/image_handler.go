package handler

import (
	"fmt"
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
		errInfos := util.MakeValidateErrInfos(err)
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	imageID := uuid.New().String()

	imageURL, err := service.UploadImageToStorage(ctx, imageID, reqBody.Base64image)
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

func GetImagesHandler(ctx *gin.Context) {
	var query dto.GetImagesQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		errInfos := util.MakeValidateErrInfos(err)
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	images, err := service.SelectImages(ctx, query.Page, query.Keyword, query.Sort, query.FavoriteImageIDs, query.AuthCheck)
	if err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

func PatchImageHandler(ctx *gin.Context) {
	imageID := ctx.Param("image_id")
	if ok := util.IsValidUUID(imageID); !ok {
		errInfos := util.MakeValidateErrInfosForUUID(imageID, "image_id")
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	var reqBody dto.PatchImageReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errInfos := util.MakeValidateErrInfos(err)
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	exists, err := service.ExistsImage(ctx, imageID)
	if err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}
	if !exists {
		msg := fmt.Sprintf("image with id: %s was not found", imageID)
		util.ResourceNotFoundResponse(ctx, msg)
		return
	}

	if err := service.UpdateImage(ctx, imageID, reqBody.Type); err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func DeleteImageHandler(ctx *gin.Context) {
	imageID := ctx.Param("image_id")
	if ok := util.IsValidUUID(imageID); !ok {
		errInfos := util.MakeValidateErrInfosForUUID(imageID, "image_id")
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	exists, err := service.ExistsImage(ctx, imageID)
	if err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}
	if !exists {
		msg := fmt.Sprintf("image with id: %s was not found", imageID)
		util.ResourceNotFoundResponse(ctx, msg)
		return
	}

	imageURL, err := service.SelectImageURL(ctx, imageID)
	if err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}

	if err := service.DeleteImage(ctx, imageID); err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}

	if err := service.DeleteImageFromStorage(imageURL); err != nil {
		util.InternalServerErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
