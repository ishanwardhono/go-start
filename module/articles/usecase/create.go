package usecase

import (
	"app/core/errors"
	handler_http "app/core/handler/http"
	"app/entity"
	"app/module/articles/repo"
	"context"
	"net/http"
	"strings"
)

type createArticle struct {
	Req         entity.Article
	RepoArticle repo.ArticleRepo
}

func (m *createArticle) Execute(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	id, err := m.RepoArticle.InsertArticle(ctx, m.Req)
	if err != nil {
		return "insert failed", err
	}
	return handler_http.Response{
		StatusCode: http.StatusCreated,
		Message:    "Success",
		Data:       id,
	}, nil
}

func (m *createArticle) Validate(ctx context.Context) error {
	var missingFields []string
	if m.Req.Title == "" {
		missingFields = append(missingFields, "title")
	}
	if m.Req.Content == "" {
		missingFields = append(missingFields, "content")
	}
	if m.Req.Author == "" {
		missingFields = append(missingFields, "author")
	}
	if len(missingFields) != 0 {
		return errors.New("missing fields: "+strings.Join(missingFields, ", "), http.StatusBadRequest)
	}
	return nil
}
