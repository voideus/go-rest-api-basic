package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/service"
)

type ArticleController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Article
	FindArticle(ctx *gin.Context)
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

func (ac *articleController) FindArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	article, err := ac.service.FindArticleById(articleId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}
