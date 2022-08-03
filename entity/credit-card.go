package entity

import "github.com/jinzhu/gorm"

type CreditCard struct {
	gorm.Model
	Number   uint   `json:"number"`
	PersonID uint64 `json:"authorID"`
}
