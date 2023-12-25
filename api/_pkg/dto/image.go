package dto

type PostImageReqBody struct {
	Image   string `json:"image" binding:"required,imagesize,base64image"`
	Keyword string `json:"keyword" binding:"max=50"`
}
