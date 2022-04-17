package repo

import (
	"app/core/database"
	"app/entity"
)

type ArticleRepo interface {
	InsertArticle(article entity.Article) error
	GetAllArticle() ([]entity.Article, error)
	GetArticle(id int) (entity.Article, error)
}

type articleRepoImpement struct {
	db *database.DB
}

func NewArticleRepo(db *database.DB) ArticleRepo {
	return &articleRepoImpement{db: db}
}

func (u *articleRepoImpement) InsertArticle(article entity.Article) error {
	_, err := u.db.NamedExec(articleInsertQuery, &article)
	return err
}

func (u *articleRepoImpement) GetAllArticle() ([]entity.Article, error) {
	articles := make([]entity.Article, 0)
	err := u.db.Select(&articles, articleGetAllQuery)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (u *articleRepoImpement) GetArticle(id int) (entity.Article, error) {
	var article entity.Article
	err := u.db.Get(&article, articleGetQuery, id)
	if err != nil {
		return entity.Article{}, err
	}
	return article, nil
}
