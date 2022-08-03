package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/service"
)

type CreditCardController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type creditCardController struct {
	service service.CreditCardService
}

func NewCreditCardController(service service.CreditCardService) CreditCardController {
	return &creditCardController{
		service: service,
	}
}

func (cc *creditCardController) Save(ctx *gin.Context) {
	var creditCard entity.CreditCard
	err := ctx.ShouldBindJSON(&creditCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	cc.service.Save(&creditCard)

	ctx.JSON(http.StatusOK, gin.H{
		"data": creditCard,
	})
}

func (cc *creditCardController) FindAll(ctx *gin.Context) {
	creditCards := cc.service.FindAll()

	ctx.JSON(http.StatusOK, gin.H{
		"creditCards": creditCards,
	})
}
