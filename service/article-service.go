package service

import (
	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/repository"
)

type ArticleService interface {
	Save(entity.Article) entity.Article
	FindAll() []entity.Article
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

func (articeService *articleService) FindAll() []entity.Article {
	return articeService.articleRepository.FindAll()

}
