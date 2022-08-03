package service

import (
	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/repository"
)

type CreditCardService interface {
	FindAll() []entity.CreditCard
	Save(creditCard *entity.CreditCard) (*entity.CreditCard, error)
}

type creditCardService struct {
	repo repository.CreditCardRepository
}

func NewCreditCardService(crRepo repository.CreditCardRepository) CreditCardService {
	return &creditCardService{
		repo: crRepo,
	}
}

func (service *creditCardService) FindAll() []entity.CreditCard {
	return service.repo.FindAll()
}

func (service *creditCardService) Save(creditCard *entity.CreditCard) (*entity.CreditCard, error) {
	return service.repo.Save(creditCard)
}
