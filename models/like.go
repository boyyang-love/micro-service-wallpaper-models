package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Like struct {
	Base
	UploadId string `form:"upload_id" json:"upload_id"`
	UserId   string `form:"user_id" json:"user_id"`
	Status   bool   `form:"status" json:"status"`
}

func (l *Like) TableName() string {
	return "like"
}

func (l *Like) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	l.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
