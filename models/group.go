package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Group struct {
	Base
	Name string `form:"name" json:"name"`
}

func (l *Group) TableName() string {
	return "group"
}

func (l *Group) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	l.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
