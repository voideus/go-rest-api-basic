package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/service"
)

type ArticleController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Article
}

type articleController struct {
	service service.ArticleService
}

func NewArticleController(service service.ArticleService) ArticleController {
	return &articleController{
		service: service,
	}
}

func (ac *articleController) FindAll() []entity.Article {
	return ac.service.FindAll()
}

func (ac *articleController) Save(ctx *gin.Context) error {
	var article entity.Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		return err
	}

	ac.service.Save(article)
	return nil
}
