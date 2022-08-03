package repository

import (
	"fmt"

	"gitlab.com/voideus/go-rest-api/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CreditCardRepository interface {
	FindAll() []entity.CreditCard
	Save(creditCard *entity.CreditCard) (*entity.CreditCard, error)
}

type creditCardRepo struct {
	db *gorm.DB
}

func NewCreditCardRepo() CreditCardRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.CreditCard{})

	return &creditCardRepo{
		db: db,
	}
}

func (crRepo *creditCardRepo) FindAll() []entity.CreditCard {
	// type Result struct {
	// 	Number    int
	// 	FirstName string
	// }
	// var results []Result

	var author entity.Person
	crRepo.db.Preload("CreditCard").Where("id = ?", 1).First(&author)

	fmt.Println("author is =>>", author)

	var creditCards []entity.CreditCard
	// err := crRepo.db.Joins("People").Find(&creditCards).Error

	// Remeber joins returns array so pass array as destination eg: [] structType
	// err := crRepo.db.Model(&entity.CreditCard{}).Select("credit_cards.number, people.first_name").Joins("left join people on people.id = credit_cards.person_id").Scan(&results).Error

	// fmt.Println(err, results)
	crRepo.db.Find(&creditCards)
	return creditCards
}

func (crRepo *creditCardRepo) Save(creditCard *entity.CreditCard) (*entity.CreditCard, error) {

	result := crRepo.db.Create(&creditCard)
	if result.Error != nil {
		err := result.Error
		return nil, err
	}

	return creditCard, nil
}
