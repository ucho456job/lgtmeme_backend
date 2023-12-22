package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var IsValidImageSize validator.Func = func(fl validator.FieldLevel) bool {
	image := fl.Field().String()
	return len(image) < 1048576*4/3
}

var IsValidBase64Image validator.Func = func(fl validator.FieldLevel) bool {
	image := fl.Field().String()
	var imagePrefixes = map[string]string{
		"jpeg": "data:image/jpeg;base64,",
		"png":  "data:image/png;base64,",
		"webp": "data:image/webp;base64,",
	}
	for _, prefix := range imagePrefixes {
		if strings.HasPrefix(image, prefix) {
			base64Data := strings.TrimPrefix(image, prefix)
			if _, err := base64.StdEncoding.DecodeString(base64Data); err == nil {
				return true
			}
		}
	}
	return false
}

func ValidateReqBody(err error) (errInfos []ErrInfo) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, valErr := range validationErrors {
			errInfos = append(errInfos, ErrInfo{
				Field:   strings.ToLower(valErr.Field()),
				Tag:     valErr.Tag(),
				Message: fmt.Sprintf("%s is invalid: %s", valErr.Field(), valErr.Error()),
			})
		}
	} else if err, isErr := err.(*json.UnmarshalTypeError); isErr {
		errInfos = append(errInfos, ErrInfo{
			Field:   err.Field,
			Tag:     "type",
			Message: fmt.Sprintf("%s is expected to be of type %s", err.Field, err.Type),
		})
	} else {
		errInfos = append(errInfos, ErrInfo{
			Field:   "",
			Tag:     "binding",
			Message: "there was an error binding the request body",
		})
	}
	return errInfos
}
