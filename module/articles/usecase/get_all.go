package usecase

import (
	"app/module/articles/repo"
	"context"
)

type getArticles struct {
	RepoArticle repo.ArticleRepo
}

func (m *getArticles) Execute(ctx context.Context) (interface{}, error) {
	m.Validate(ctx)
	return m.RepoArticle.GetAllArticle()
}

func (m *getArticles) Validate(ctx context.Context) error {
	return nil
}
