package repo

import (
	"app/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	InsertUser(user entity.User) error
	GetAllUser() ([]entity.User, error)
	GetUser(name string) (entity.User, error)
}

type userRepoImpement struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepoImpement{db: db}
}

func (u *userRepoImpement) InsertUser(user entity.User) error {
	err := u.db.Get(&user, userInsertQuery)
	return err
}

func (u *userRepoImpement) GetAllUser() ([]entity.User, error) {
	users := make([]entity.User, 0)
	err := u.db.Select(&users, userGetAllQuery)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepoImpement) GetUser(name string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(userGetQuery, name).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
