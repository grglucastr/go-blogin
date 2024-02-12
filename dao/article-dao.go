package dao

import (
	"database/sql"

	"github.com/grglucastr/go-blogin/models"
)

type ArticleDAO struct {
	db *sql.DB
}

func NewArticleDAO(db *sql.DB) *ArticleDAO {
	return &ArticleDAO{
		db: db,
	}
}

func (ad *ArticleDAO) Save(article *models.Article) (*models.Article, error) {
	sql := "INSERT INTO articles (id, title, content) VALUES (?, ?, ?)"

	_, err := ad.db.Exec(sql, article.ID, article.Title, article.Content)

	if err != nil {
		return nil, err
	}

	return article, nil
}
