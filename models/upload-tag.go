package models

type UploadTag struct {
	UploadId string `form:"upload_id" json:"upload_id"`
	TagId    string `form:"tag_id" json:"tag_id"`
}

func (u *UploadTag) TableName() string {
	return "upload_tag"
}
