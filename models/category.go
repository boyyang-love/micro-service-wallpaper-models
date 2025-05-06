package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Category struct {
	Base
	Name string `json:"name" form:"name"`
	Web  bool   `json:"web"`
	Moa  bool   `json:"moa"`
	Sort int    `json:"sort" form:"sort"`
}

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	c.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
