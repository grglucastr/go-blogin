package controllers

import (
	"encoding/json"
	"net/http"

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


func (ac *ArticleController) PostArticle(r *http.Request, w http.ResponseWriter){
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