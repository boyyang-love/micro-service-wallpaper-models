package models

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Post struct {
	Base
	UserId       string `json:"user_id" form:"user_id"`
	Title        string `json:"title" form:"title"`
	Content      string `json:"content" form:"content"`
	ImageIds     string `json:"image_ids" form:"image_ids"`            // 逗号分隔的 upload id
	Status       int    `json:"status" form:"status" gorm:"default:1"` // 0=草稿 1=已发布 2=审核中
	RejectReason string `json:"reject_reason" form:"reject_reason" gorm:"type:text"`
}

func (p *Post) TableName() string {
	return "post"
}

func (p *Post) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	p.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
