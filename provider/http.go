package provider

import (
	"app/core/handler/http"
	articles "app/module/articles/handler/http"
)

// list of handlers
func GetHttpHandlers() []http.HttpHandler {
	return []http.HttpHandler{
		getArticleHandler(),
	}
}

// create handler instance
func getArticleHandler() http.HttpHandler {
	return articles.NewArticleHandler(
		GetArticleFactory(),
	)
}
