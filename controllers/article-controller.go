package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grglucastr/go-blogin/dao"
	"github.com/grglucastr/go-blogin/models"
)

type ArticleController struct {
	articleDAO *dao.ArticleDAO
}

func NewArticleController(articleDAO *dao.ArticleDAO) *ArticleController {
	return &ArticleController{
		articleDAO: articleDAO,
	}
}

func (ac *ArticleController) PostArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ac.articleDAO.Save(&article)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (ac *ArticleController) ListArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := ac.articleDAO.ListAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articles)
}

func (ac *ArticleController) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	article, err := ac.articleDAO.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}
