package repo

import (
	"app/core/database"
	"app/module/articles/model"
	"context"
)

type ArticleRepo interface {
	InsertArticle(ctx context.Context, article model.Article) (int, error)
	GetAllArticle(ctx context.Context) ([]model.Article, error)
	GetArticle(ctx context.Context, id int) (model.Article, error)
}

type articleRepoImpement struct {
	db *database.DB
}

func NewArticleRepo(db *database.DB) ArticleRepo {
	return &articleRepoImpement{db: db}
}

func (u *articleRepoImpement) InsertArticle(ctx context.Context, article model.Article) (int, error) {
	var id int
	rows, err := u.db.NamedQueryContext(ctx, articleInsertQuery, &article)
	rows.Next()
	rows.Scan(&id)
	return id, err
}

func (u *articleRepoImpement) GetAllArticle(ctx context.Context) ([]model.Article, error) {
	articles := make([]model.Article, 0)
	err := u.db.SelectContext(ctx, &articles, articleGetAllQuery)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (u *articleRepoImpement) GetArticle(ctx context.Context, id int) (model.Article, error) {
	var article model.Article
	err := u.db.GetContext(ctx, &article, articleGetQuery, id)
	if err != nil {
		return model.Article{}, err
	}
	return article, nil
}
