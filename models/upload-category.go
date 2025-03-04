package models

type UploadCategory struct {
	UploadId   string `form:"upload_id" json:"upload_id"`
	CategoryId string `form:"category_id" json:"category_id"`
}

func (u *UploadCategory) TableName() string {
	return "upload_category"
}
