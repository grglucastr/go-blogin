package models

import (
	"fmt"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}


func NewPost(id string, title string, content string) *Post {
	return &Post{
		ID:      id,
		Title:   title,
		Content: content,
	}
}

func (p *Post) DoSomething(action string) {
	fmt.Println("doing something like: ", action)
}

func (p *Post) ChangeTitle(newTitle string){
	p.Title = newTitle
}