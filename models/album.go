package models

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Album struct {
	Base
	Name   string `json:"name" form:"name"`
	Desc   string `json:"desc" form:"desc"`
	Status bool   `json:"status" form:"status"`
}

func (a *Album) TableName() string {
	return "album"
}

func (a *Album) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	a.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
