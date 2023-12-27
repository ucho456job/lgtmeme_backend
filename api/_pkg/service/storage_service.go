package service

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{}

func UploadImageToStorage(ctx *gin.Context, ID string, base64Image string) (imageURL string, err error) {
	contentType, err := getContentType(base64Image)
	if err != nil {
		return "", err
	}

	decodedData, err := decodeBase64Image(base64Image)
	if err != nil {
		return "", err
	}

	uploadURL := buildStorageURL(ID, contentType, false)
	if err := uploadToStorage(uploadURL, decodedData, contentType); err != nil {
		return "", err
	}

	imageURL = buildStorageURL(ID, contentType, true)
	return imageURL, nil
}

func DeleteImageFromStorage(imageURL string) error {
	uploadURL := strings.Replace(imageURL, "/public", "", 1)
	return makeStorageRequest("DELETE", uploadURL, nil, "")
}

func getContentType(base64Image string) (string, error) {
	switch {
	case strings.HasPrefix(base64Image, "data:image/jpeg;base64,"):
		return "image/jpeg", nil
	case strings.HasPrefix(base64Image, "data:image/png;base64,"):
		return "image/png", nil
	case strings.HasPrefix(base64Image, "data:image/webp;base64,"):
		return "image/webp", nil
	}
	return "", errors.New("unsupported image type")
}

func decodeBase64Image(base64Image string) ([]byte, error) {
	dataParts := strings.SplitN(base64Image, ",", 2)
	if len(dataParts) != 2 {
		return nil, errors.New("invalid base64 format")
	}
	return base64.StdEncoding.DecodeString(dataParts[1])
}

func buildStorageURL(ID, contentType string, public bool) string {
	extension := map[string]string{
		"image/jpeg": "jpg",
		"image/png":  "png",
		"image/webp": "webp",
	}[contentType]

	baseUrl := os.Getenv("SUPABASE_URL")
	path := "storage/v1/object"
	if public {
		path += "/public"
	}
	return fmt.Sprintf("%s/%s/images/%s.%s", baseUrl, path, ID, extension)
}

func uploadToStorage(url string, data []byte, contentType string) error {
	return makeStorageRequest("POST", url, bytes.NewReader(data), contentType)
}

func makeStorageRequest(method, url string, body io.Reader, contentType string) error {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	anonKey := os.Getenv("SUPABASE_ANON_KEY")
	request.Header.Set("Authorization", "Bearer "+anonKey)
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("make request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("read response body: %w", err)
		}
		return fmt.Errorf("request failed: %s", string(bodyBytes))
	}

	return nil
}
