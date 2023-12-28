package dto

type PostImageReqBody struct {
	Base64image string `json:"base64image" binding:"required,imagesize,base64image"`
	Keyword     string `json:"keyword" binding:"max=50"`
}

type GetImageSort string

const (
	GetImageSortLatest  GetImageSort = "latest"
	GetImageSortPopular GetImageSort = "popular"
)

type GetImagesQuery struct {
	Page             int          `form:"page" binding:"min=0"`
	Keyword          string       `form:"keyword" binding:"max=50"`
	Sort             GetImageSort `form:"sort" binding:"omitempty"`
	FavoriteImageIDs []string     `form:"favorite_image_ids" binding:"omitempty,uuidslice"`
	AuthCheck        bool         `form:"auth_check" binding:"omitempty"`
}

type PatchImageReqType string

const (
	PatchImageReqTypeUsed      PatchImageReqType = "used"
	PatchImageReqTypeReporting PatchImageReqType = "reporting"
	PatchImageReqTypeConfirmed PatchImageReqType = "confirmed"
)

type PatchImageReqBody struct {
	Type PatchImageReqType `json:"type" binding:"required"`
}
