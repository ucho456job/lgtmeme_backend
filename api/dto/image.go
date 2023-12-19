package dto

type PostImageReqBody struct {
	Image   string `json:"image" binding:"required"`
	Keyword string `json:"keyword" binding:"max=50"`
}
