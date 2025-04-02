package models

type UploadGroup struct {
	UploadId string `form:"upload_id" json:"upload_id"`
	GroupId  string `form:"group_id" json:"group_id"`
}

func (u *UploadGroup) TableName() string {
	return "upload_group"
}
