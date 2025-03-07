package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Recommend struct {
	Base
	Name string `json:"name" form:"name"`
	Sort int    `json:"sort" form:"sort"`
}

func (r *Recommend) TableName() string {
	return "recommend"
}

func (r *Recommend) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	r.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
