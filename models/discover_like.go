package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type DiscoverLike struct {
	Base
	UserId     string `json:"user_id" form:"user_id"`
	DiscoverId string `json:"discover_id" form:"discover_id"`
}

func (dl *DiscoverLike) TableName() string {
	return "discover_likes"
}

func (dl *DiscoverLike) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	dl.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
