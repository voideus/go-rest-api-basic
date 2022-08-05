package entity

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title       string `json:"title" binding:"required" gorm:"type:varchar(32)"`
	Description string `json:"description" binding:"required" gorm:"type:varchar(32)"`
	Body        string `json:"body" binding:"required" gorm:"type:varchar(255)"`
	PersonID    uint64 `json:"authorId"`
	Person      Person ` gorm:"foreignKey:PersonID"`
	Comments    []Comment
}

type Comment struct {
	Comment   string
	ArticleID uint
}

type CommentInput struct {
	Comment string `json:"comment" binding:"required"`
}
