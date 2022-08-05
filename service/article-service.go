package service

import (
	"strconv"

	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/repository"
)

type ArticleService interface {
	Save(entity.Article) entity.Article
	FindAll() []entity.Article
	FindArticleById(articleId string) (*entity.Article, error)
	AddComment(articleId string, comment string) (*entity.Comment, error)
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleRepoService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleService{
		articleRepository: articleRepo,
	}
}

func (articeService *articleService) Save(article entity.Article) entity.Article {
	articeService.articleRepository.Save(article)
	return article
}

func (articeService *articleService) AddComment(articleId string, comment string) (*entity.Comment, error) {
	articleID, err := strconv.Atoi(articleId)
	if err != nil {
		return nil, err
	}
	commentToSave := entity.Comment{Comment: comment, ArticleID: uint(articleID)}
	return articeService.articleRepository.AddComment(commentToSave)
}

func (articeService *articleService) FindAll() []entity.Article {
	return articeService.articleRepository.FindAll()
}

func (articeService *articleService) FindArticleById(articleId string) (*entity.Article, error) {
	id, err := strconv.Atoi(articleId)
	if err != nil {
		return nil, err
	}

	return articeService.articleRepository.FindById(id)
}
