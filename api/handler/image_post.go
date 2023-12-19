package handler

import (
	"lgtmeme_backend/api/config"
	"lgtmeme_backend/api/dto"
	"lgtmeme_backend/api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostImageHandler(ctx *gin.Context) {
	if errInfos, ok := validatePostImageReqBody(ctx); !ok {
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	ctx.String(http.StatusCreated, "Success")
}

func validatePostImageReqBody(ctx *gin.Context) (errInfos []util.ErrInfo, ok bool) {
	ok = true
	var reqBody dto.PostImageReqBody
	var customErrMsgs = map[string]map[string]string{
		"Image": {
			"required": "image is required.",
			"base64":   "image must be a valid base64 string.",
			"size":     "image size must be less than 1MB.",
		},
		"Keyword": {
			"max": "keyword should be less than 50 characters.",
		},
	}

	errInfos, ok = util.ValidateStruct(ctx, &reqBody, customErrMsgs)

	if !util.ValidateBase64Image(ctx, reqBody.Image) {
		errInfos = append(errInfos, util.ErrInfo{
			Field:   "image",
			Tag:     "format",
			Message: customErrMsgs["Image"]["base64"],
		})
		ok = false
	}

	if len(reqBody.Image) > config.MAX_IMAGE_SIZE {
		errInfos = append(errInfos, util.ErrInfo{
			Field:   "image",
			Tag:     "size",
			Message: customErrMsgs["Image"]["size"],
		})
		ok = false
	}

	return errInfos, ok
}
