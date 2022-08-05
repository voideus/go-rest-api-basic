package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/service"
)

type PersonController interface {
	FindAll(ctx *gin.Context)
	AddLanguagesToPerson(ctx *gin.Context)
	GetLanguagesOfPerson(ctx *gin.Context)
}

type personController struct {
	service service.PersonService
}

func NewPersonController(service service.PersonService) PersonController {
	return &personController{
		service: service,
	}
}

func (pc *personController) FindAll(ctx *gin.Context) {
	people := []entity.Person{}
	people = pc.service.FindAll()

	ctx.JSON(http.StatusOK, gin.H{
		"people": people,
	})

}

func (pc *personController) AddLanguagesToPerson(ctx *gin.Context) {
	var languageInput entity.LanguageInput
	err := ctx.ShouldBindJSON(&languageInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error binding": err.Error()})
	}

	err = pc.service.AddLanguageToPerson(ctx.Param("id"), languageInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Languages added successfully",
	})
}

func (pc *personController) GetLanguagesOfPerson(ctx *gin.Context) {
	person, err := pc.service.GetLanguagesOfPerson(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": person,
	})
}
