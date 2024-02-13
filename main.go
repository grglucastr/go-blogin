package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grglucastr/go-blogin/controllers"
	"github.com/grglucastr/go-blogin/dao"
)

func main() {
	db, err := sql.Open("mysql", "root:thepass@tcp(localhost:3306)/go-blogin")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	articleDAO := dao.NewArticleDAO(db)
	articleController := controllers.NewArticleController(articleDAO)

	c := chi.NewRouter()

	c.Post("/articles", articleController.PostArticle)
	c.Get("/articles", articleController.ListArticles)
	c.Get("/articles/{id}", articleController.GetById)
	c.Delete("/articles/{id}", articleController.DeleteById)

	fmt.Println("Server is runnint on port 8080")

	http.ListenAndServe(":8080", c)

}
