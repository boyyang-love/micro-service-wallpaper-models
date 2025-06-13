package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Block struct {
	Base
	TargetId string `json:"target_id" form:"target_id"`
	Type     string `json:"type" form:"type"`
	UserId   string `json:"user_id" form:"user_id"`
}

func (b *Block) TableName() string {
	return "block"
}

func (b *Block) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	b.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
