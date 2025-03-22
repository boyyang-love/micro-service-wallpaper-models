package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Carousel struct {
	Base
	Path string `json:"name" form:"name"`
	Sort int    `json:"sort" form:"sort"`
}

func (c *Carousel) TableName() string {
	return "carousel"
}

func (c *Carousel) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	c.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
