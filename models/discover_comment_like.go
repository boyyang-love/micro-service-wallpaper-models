package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type DiscoverCommentLike struct {
	Base
	UserId    string `json:"user_id" form:"user_id"`
	CommentId string `json:"comment_id" form:"comment_id"`
}

func (dcl *DiscoverCommentLike) TableName() string {
	return "discover_comment_likes"
}

func (dcl *DiscoverCommentLike) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.NewV1()
	dcl.Id = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
