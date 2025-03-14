package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Download struct {
	Base
	DownloadId string `json:"download_id" form:"download_id"`
	UserId     string `json:"user_id" form:"user_id"`
}

func (d *Download) TableName() string {
	return "download"
}

func (d *Download) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	d.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
