package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Search struct {
	Base
	Keywords    string `json:"keywords" form:"keywords"`
	SearchCount int64  `json:"search_count" form:"search_count"`
}

func (s *Search) TableName() string {
	return "search"
}

func (s *Search) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	s.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
