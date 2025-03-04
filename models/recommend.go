package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Recommend struct {
	Base
	Name string `json:"name" form:"name"`
}

func (r *Recommend) TableName() string {
	return "like"
}

func (r *Recommend) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	r.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
