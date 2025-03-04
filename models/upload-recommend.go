package models

type UploadRecommend struct {
	UploadId    string `form:"upload_id" json:"upload_id"`
	RecommendId string `form:"recommend_id" json:"recommend_id"`
}

func (u *UploadRecommend) TableName() string {
	return "upload_recommend"
}
