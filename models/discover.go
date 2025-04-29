package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Discover struct {
	Base
	Title    string `json:"title" form:"title"`
	Subtitle string `json:"subtitle" form:"subtitle"`
	UserId   string `json:"user_id" form:"user_id"`
	ImageIds string `json:"image_ids" form:"image_ids"`
	Status   int    `json:"status" form:"status" gorm:"default:1"`
}

func (d *Discover) TableName() string {
	return "discover"
}

func (d *Discover) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	d.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
