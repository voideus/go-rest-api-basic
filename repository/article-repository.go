package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.com/voideus/go-rest-api/entity"
)

type ArticleRepository interface {
	Save(article entity.Article)
	FindAll() []entity.Article
	FindOne(id string) entity.Article
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	db, err := gorm.Open("sqlite3", "test.db")
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

func (ar *articleRepository) FindOne(id string) entity.Article {
	var article entity.Article
	ar.db.Preload("Person").Where("id = ?", id).First(&article)
	return article
}
