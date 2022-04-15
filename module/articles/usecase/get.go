package usecase

import (
	"app/core/errors"
	"app/module/articles/repo"
	"context"
	"net/http"
)

type getArticle struct {
	Req         int
	RepoArticle repo.ArticleRepo
}

func (m *getArticle) Execute(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	return m.RepoArticle.GetArticle(m.Req)
}

func (m *getArticle) Validate(ctx context.Context) error {
	if m.Req == 0 {
		return errors.New("missing field id", http.StatusBadRequest)
	}
	return nil
}
