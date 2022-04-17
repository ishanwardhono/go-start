package handler

import (
	http_handler "app/core/handler/http"
	"app/core/log"
	"app/entity"
	articles "app/module/articles/usecase"
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	articles articles.Factory
}

func NewArticleHandler(articles articles.Factory) http_handler.HttpHandler {
	return &ArticleHandler{
		articles: articles,
	}
}

func (uh *ArticleHandler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/articles", http_handler.Handle(uh.allArticles)).Methods("GET")
	router.HandleFunc("/article/{id}", http_handler.Handle(uh.getArticle)).Methods("GET")
	router.HandleFunc("/article", http_handler.Handle(uh.newArticle)).Methods("POST")
	router.HandleFunc("/", http_handler.Handle(func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return "test endpoint", nil
	}))
}

func (uh *ArticleHandler) allArticles(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return uh.articles.GetAll().Execute(ctx)
}

func (uh *ArticleHandler) newArticle(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var req entity.Article

	err := http_handler.ParseBody(r, &req)
	if err != nil {
		log.Error(ctx, "Failed parse request")
		return nil, err
	}

	return uh.articles.Create(req).Execute(ctx)
}

func (uh *ArticleHandler) getArticle(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idStr := http_handler.GetQueryParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(ctx, "id is not a number")
		return nil, err
	}
	return uh.articles.Get(id).Execute(ctx)
}
