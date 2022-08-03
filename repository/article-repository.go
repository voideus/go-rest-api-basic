package repository

import (
	"gitlab.com/voideus/go-rest-api/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Save(article entity.Article)
	FindAll() []entity.Article
	FindById(id int) (*entity.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Article{})

	return &articleRepository{
		db: db,
	}
}

func (ar *articleRepository) Save(article entity.Article) {
	ar.db.Create(&article)
}

func (ar *articleRepository) FindAll() []entity.Article {
	var articles []entity.Article
	ar.db.Preload("Person").Find(&articles)

	return articles
}

func (ar *articleRepository) FindById(id int) (*entity.Article, error) {
	var article entity.Article
	result := ar.db.Preload("Person").Where("id = ?", id).First(&article)

	if result.Error != nil {
		err := result.Error

		return nil, err
	}

	return &article, nil
}
