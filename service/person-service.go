package service

import (
	"strconv"

	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/repository"
)

type PersonService interface {
	FindAll() []entity.Person
	AddLanguageToPerson(personID string, languageInput entity.LanguageInput) error
	GetLanguagesOfPerson(personID string) (*entity.Person, error)
}

type personService struct {
	repo repository.PersonRepo
}

func NewPersonService(pr repository.PersonRepo) PersonService {
	return &personService{
		repo: pr,
	}
}

func (ps *personService) FindAll() []entity.Person {
	return ps.repo.FindAll()
}

func (ps *personService) AddLanguageToPerson(personID string, languageInput entity.LanguageInput) error {
	personId, err := strconv.Atoi(personID)

	if err != nil {
		return err
	}
	var languages []entity.Language

	for _, e := range languageInput.Languages {
		languages = append(languages, entity.Language{
			Name: e.Name,
		})
	}

	return ps.repo.AddLanguageToPerson(uint(personId), languages)
}

func (ps *personService) GetLanguagesOfPerson(personID string) (*entity.Person, error) {
	personId, err := strconv.ParseUint(personID, 0, 0)

	if err != nil {
		return nil, err
	}

	return ps.repo.GetLanguagesOfPerson(uint(personId))
}
