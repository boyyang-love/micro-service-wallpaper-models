package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type UserFollow struct {
	Base
	FollowerId  string `json:"follower_id" form:"follower_id"`
	FollowingId string `json:"following_id" form:"following_id"`
}

func (uf *UserFollow) TableName() string {
	return "user_follows"
}

func (uf *UserFollow) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	uf.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
