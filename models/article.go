package models

import (
	"fmt"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewArticle(id string, title string, content string) *Article {
	return &Article{
		ID:      id,
		Title:   title,
		Content: content,
	}
}

func (a *Article) DoSomething(action string) {
	fmt.Println("doing something like: ", action)
}

func (a *Article) ChangeTitle(newTitle string) {
	a.Title = newTitle
}
