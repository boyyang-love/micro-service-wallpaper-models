package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Upload struct {
	Base
	Hash           string `json:"hash"`
	FileName       string `json:"file_name" form:"file_name"`
	OriginFileSize int64  `json:"origin_file_size" form:"origin_file_size"`
	FileSize       int64  `json:"file_size" form:"file_size"`
	OriginType     string `json:"origin_type" form:"origin_type"`
	FileType       string `json:"file_type" form:"file_type"`
	OriginFilePath string `json:"origin_file_path" form:"origin_file_path"`
	FilePath       string `json:"file_path" form:"file_path"`
	Type           string `json:"type" form:"type"`
	W              int    `json:"w" form:"w"`
	H              int    `json:"h" form:"h"`
	Status         int    `json:"status" form:"status" gorm:"default:1"`
	Download       int    `json:"download" form:"download" gorm:"default:0"`
	View           int    `json:"view" form:"view" gorm:"default:0"`
	UserId         string `json:"user_id" form:"user_id"`
}

// 定义Upload结构体的TableName方法，返回表名为"upload"
func (u *Upload) TableName() string {
	return "upload"
}

func (u *Upload) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	u.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
