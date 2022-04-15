package provider

import (
	"app/core/handler/http"
	articles "app/module/articles/handler"
)

//list of handlers
func GetHandlers() []http.HttpHandler {
	return []http.HttpHandler{
		GetArticleHandler(),
	}
}

//create handler instance
func GetArticleHandler() http.HttpHandler {
	return articles.NewArticleHandler(
		GetArticleFactory(),
	)
}
