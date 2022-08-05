package repository

import (
	"gitlab.com/voideus/go-rest-api/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PersonRepo interface {
	FindAll() []entity.Person
	AddLanguageToPerson(personID uint, languages []entity.Language) error
	GetLanguagesOfPerson(personID uint) (*entity.Person, error)
}

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo() PersonRepo {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Language{})

	return &personRepo{
		db: db,
	}
}

func (pr *personRepo) FindAll() []entity.Person {
	var people []entity.Person
	pr.db.Find(&people)
	return people
}

func (pr *personRepo) AddLanguageToPerson(personID uint, languages []entity.Language) error {
	var person entity.Person
	err := pr.db.Model(&entity.Person{}).Where("id = ?", personID).First(&person).Error
	if err != nil {
		return err
	}

	return pr.db.Model(&person).Association("Languages").Append(languages)
}

func (pr *personRepo) GetLanguagesOfPerson(personID uint) (*entity.Person, error) {
	var person = entity.Person{}
	return &person, pr.db.Model(&person).Where("id = ?", personID).Preload("Languages").First(&person).Error
}
