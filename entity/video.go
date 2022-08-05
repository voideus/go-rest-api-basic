package entity

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	FirstName  string `json:"firstname"  gorm:"type:varchar(32)"`
	LastName   string `json:"lastname"  gorm:"type:varchar(32)"`
	Age        int8   `json:"age" `
	Email      string `json:"email" gorm:"type:varchar(256)"`
	CreditCard *CreditCard
	Languages  []*Language `gorm:"many2many:user_languages;"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=100" validate:"is-cool" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP" `
}

type Language struct {
	gorm.Model
	Name   string
	People []*Person `gorm:"many2many:user_languages;"`
}

type LanguageOnly struct {
	Name string `json:"name"`
}

type LanguageInput struct {
	Languages []LanguageOnly `json:"languages"`
}
