package provider

import (
	"app/core/database"
	articles "app/module/articles/repo"
)

//create repo instance
func GetArticlesRepo() articles.ArticleRepo {
	return articles.NewArticleRepo(
		database.GetDB(),
	)
}
