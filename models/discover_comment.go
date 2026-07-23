package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type DiscoverComment struct {
	Base
	UserId     string `json:"user_id" form:"user_id"`
	DiscoverId string `json:"discover_id" form:"discover_id"`
	ParentId   string `json:"parent_id" form:"parent_id"`
	Content    string `json:"content" form:"content"`
	LikeCount  int    `json:"like_count" form:"like_count" gorm:"default:0"`
	ReplyCount int    `json:"reply_count" form:"reply_count" gorm:"default:0"`
	Status     int    `json:"status" form:"status" gorm:"default:1"` // 1:正常, 0:删除
}

func (dc *DiscoverComment) TableName() string {
	return "discover_comments"
}

func (dc *DiscoverComment) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	dc.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
