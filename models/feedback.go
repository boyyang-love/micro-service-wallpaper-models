package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Feedback struct {
	Base
	FeedbackId string `json:"feedback_id" form:"feedback_id"`
	Content    string `json:"Content" form:"Content"`
	Type       string `json:"type" form:"type"`
	Status     int    `json:"status" form:"status"`
	UserId     string `json:"user_id" form:"user_id"`
}

func (f *Feedback) TableName() string {
	return "feedback"
}

func (f *Feedback) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	f.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
