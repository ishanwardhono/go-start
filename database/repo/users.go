package repo

import (
	"database/sql"
	"sm-secret/entity"
)

type UserRepo interface {
	InsertUser(user entity.User) error
	GetAllUser() ([]entity.User, error)
	GetUser(name string) (entity.User, error)
}

type userRepoImpement struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpement{db: db}
}

func (u *userRepoImpement) InsertUser(user entity.User) error {
	_, err := u.db.Exec(userInsertQuery, user.Name, user.Email)
	return err
}

func (u *userRepoImpement) GetAllUser() ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := u.db.Query(userGetAllQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepoImpement) GetUser(name string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(userGetQuery, name).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return entity.User{}, nil
	}
	return user, nil
}
