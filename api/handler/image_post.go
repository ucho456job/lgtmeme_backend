package handler

import (
	"lgtmeme_backend/api/dto"
	"lgtmeme_backend/api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostImageHandler(ctx *gin.Context) {
	var reqBody dto.PostImageReqBody

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errInfos := util.ValidateReqBody(ctx, err)
		util.ValidationFailedResponse(ctx, errInfos)
		return
	}

	ctx.String(http.StatusCreated, "Success")
}
