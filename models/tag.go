package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Tag struct {
	Base
	// 字段
	Name   string `json:"name" form:"tag_name"`
	Type   string `json:"type" form:"type"`
	Sort   int    `json:"sort" form:"sort"`
	UserId string `json:"user_id" form:"user_id"`
}

func (t *Tag) TableName() string {
	return "tag"
}

func (t *Tag) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	t.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
