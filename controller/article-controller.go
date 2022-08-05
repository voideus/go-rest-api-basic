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
	AddCommentToArticle(ctx *gin.Context)
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

func (ac *articleController) AddCommentToArticle(ctx *gin.Context) {
	var comment entity.CommentInput
	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"should bind error": err.Error(),
		})
		return
	}

	err = validate.Struct(comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"validate struct error": err.Error(),
		})
		return
	}

	articleId := ctx.Param("id")

	_, err1 := ac.service.FindArticleById(articleId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err1.Error(),
		})
		return
	}

	_, err2 := ac.service.AddComment(articleId, comment.Comment)

	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err2.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "comment added successfully",
	})
}
