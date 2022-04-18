package provider

import articles "app/module/articles/business"

//set module factory & inject dependency
func GetArticleFactory() articles.Factory {
	return articles.NewArticlesFactory(
		GetArticlesRepo(),
	)
}
