package dao

import (
	"database/sql"

	"github.com/grglucastr/go-blogin/models"
)

type PostDAO struct {
	db *sql.DB
}

func NewPostDAO(db *sql.DB) *PostDAO {
	return &PostDAO{
		db: db,
	}
}

func (pd *PostDAO) Save(post *models.Post) (*models.Post, error) {
	sql := "INSERT INTO posts (id, title, content) VALUES (?, ?, ?)"

	_, err := pd.db.Exec(sql, post.ID, post.Title, post.Content)

	if err != nil {
		return nil, err
	}

	return post, nil
}
