package models

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Base
	PostId       string `json:"post_id" form:"post_id"`
	UserId       string `json:"user_id" form:"user_id"`
	Content      string `json:"content" form:"content"`
	Status       int    `json:"status" form:"status" gorm:"default:1"` // 0=删除 1=正常
	RejectReason string `json:"reject_reason" form:"reject_reason" gorm:"type:text"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	c.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
