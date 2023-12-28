package dto

type PostImageReqBody struct {
	Base64image string `json:"base64image" binding:"required,imagesize,base64image"`
	Keyword     string `json:"keyword" binding:"max=50"`
}

type GetImagesQuery struct {
	Page        int      `form:"page" binding:"min=0"`
	Keyword     string   `form:"keyword" binding:"max=50"`
	Sort        string   `form:"sort" binding:"omitempty,oneof=latest popular"`
	FavoriteIDs []string `form:"favorite_ids" binding:"omitempty,uuidslice"`
	AuthCheck   bool     `form:"auth_check" binding:"omitempty"`
}

type PatchImageReqBody struct {
	Type string `json:"type" binding:"required,oneof=used reporting confirmed"`
}
