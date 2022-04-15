package provider

import articles "app/module/articles/usecase"

//set module factory & inject dependency
func GetArticleFactory() articles.Factory {
	return articles.NewArticlesFactory(
		GetArticlesRepo(),
	)
}
