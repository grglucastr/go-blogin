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

func (ad *ArticleDAO) ListAll() ([]*models.Article, error) {
	rows, err := ad.db.Query("SELECT id, title, content FROM articles")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []*models.Article

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}

		articles = append(articles, &article)
	}

	return articles, nil
}

func (ad *ArticleDAO) GetById(uuid string) (*models.Article, error) {

	var article models.Article
	var sql string = "SELECT id, title, content FROM articles WHERE id = ?"

	err := ad.db.QueryRow(sql, uuid).Scan(&article.ID, &article.Title, &article.Content)

	if err != nil {
		return nil, err
	}

	return &article, nil
}


func (ad *ArticleDAO) DeleteById(uuid string) (error) {

	_, err := ad.db.Exec("DELETE FROM articles WHERE id = ?", uuid)
	return err
}