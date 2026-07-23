package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type DiscoverFavorite struct {
	Base
	UserId     string `json:"user_id" form:"user_id"`
	DiscoverId string `json:"discover_id" form:"discover_id"`
}

func (df *DiscoverFavorite) TableName() string {
	return "discover_favorites"
}

func (df *DiscoverFavorite) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	df.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
