package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(ctx *gin.Context, s interface{}, customErrMsgs map[string]map[string]string) (errInfos []ErrInfo, ok bool) {
	if err := ctx.ShouldBindJSON(&s); err != nil {
		if e, isErr := err.(*json.UnmarshalTypeError); isErr {
			return []ErrInfo{{
				Field:   e.Field,
				Tag:     "type",
				Message: fmt.Sprintf("%s is invalid type.", e.Field),
			}}, false
		}

		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fieldErr := customErrMsgs[e.Field()]
			msg := "Unknown error occurs."
			if customMsg, exists := fieldErr[e.Tag()]; exists {
				msg = customMsg
			}
			errInfos = append(errInfos, ErrInfo{
				Field:   e.Field(),
				Tag:     e.Tag(),
				Message: msg,
			})
		}
		return errInfos, false
	}
	return nil, true
}

func ValidateBase64Image(ctx *gin.Context, imageData string) (ok bool) {
	var imagePrefixes = map[string]string{
		"jpeg": "data:image/jpeg;base64,",
		"png":  "data:image/png;base64,",
		"webp": "data:image/webp;base64,",
	}
	for _, prefix := range imagePrefixes {
		if strings.HasPrefix(imageData, prefix) {
			base64Data := strings.TrimPrefix(imageData, prefix)
			if _, err := base64.StdEncoding.DecodeString(base64Data); err != nil {
				return false
			}
			return true
		}
	}
	return false
}
