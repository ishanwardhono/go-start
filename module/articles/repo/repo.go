package repo

import (
	"app/core/database"
	"app/entity"
	"context"
)

type ArticleRepo interface {
	InsertArticle(ctx context.Context, article entity.Article) (int, error)
	GetAllArticle(ctx context.Context) ([]entity.Article, error)
	GetArticle(ctx context.Context, id int) (entity.Article, error)
}

type articleRepoImpement struct {
	db *database.DB
}

func NewArticleRepo(db *database.DB) ArticleRepo {
	return &articleRepoImpement{db: db}
}

func (u *articleRepoImpement) InsertArticle(ctx context.Context, article entity.Article) (int, error) {
	var id int
	rows, err := u.db.NamedQueryContext(ctx, articleInsertQuery, &article)
	rows.Next()
	rows.Scan(&id)
	return id, err
}

func (u *articleRepoImpement) GetAllArticle(ctx context.Context) ([]entity.Article, error) {
	articles := make([]entity.Article, 0)
	err := u.db.SelectContext(ctx, &articles, articleGetAllQuery)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (u *articleRepoImpement) GetArticle(ctx context.Context, id int) (entity.Article, error) {
	var article entity.Article
	err := u.db.GetContext(ctx, &article, articleGetQuery, id)
	if err != nil {
		return entity.Article{}, err
	}
	return article, nil
}
