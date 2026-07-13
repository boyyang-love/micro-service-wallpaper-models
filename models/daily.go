package models

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DailyWallpaper struct {
	Base
	UploadId    string `json:"upload_id" form:"upload_id" gorm:"column:upload_id"`
	Date        string `json:"date" form:"date" gorm:"column:date;type:varchar(10);uniqueIndex"`
	Description string `json:"description" form:"description" gorm:"column:description"`
	Edition     int    `json:"edition" form:"edition" gorm:"column:edition"`
	Status      int    `json:"status" form:"status" gorm:"column:status;default:1"`
}

func (d *DailyWallpaper) TableName() string {
	return "daily_wallpaper"
}

func (d *DailyWallpaper) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	d.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
