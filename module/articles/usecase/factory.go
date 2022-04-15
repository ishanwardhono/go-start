package usecase

import (
	"app/entity"
	"app/module"
	"app/module/articles/repo"
)

type Factory interface {
	Create(req entity.Article) module.BaseModule
	Get(id int) module.BaseModule
	GetAll() module.BaseModule
}

type articlesFactory struct {
	repoArticle repo.ArticleRepo
}

func NewArticlesFactory(repoArticle repo.ArticleRepo) Factory {
	return &articlesFactory{
		repoArticle: repoArticle,
	}
}

func (f *articlesFactory) Create(req entity.Article) module.BaseModule {
	return &createArticle{
		Req:         req,
		RepoArticle: f.repoArticle,
	}
}

func (f *articlesFactory) Get(id int) module.BaseModule {
	return &getArticle{
		Req:         id,
		RepoArticle: f.repoArticle,
	}
}

func (f *articlesFactory) GetAll() module.BaseModule {
	return &getArticles{
		RepoArticle: f.repoArticle,
	}
}
