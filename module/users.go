package module

import (
	"app/entity"
	"app/provider/database/repo"
)

type UserModel struct {
	repo repo.UserRepo
}

func NewUserModel(repo repo.UserRepo) UserModel {
	return UserModel{repo: repo}
}

func (m *UserModel) InsertUser(user entity.User) error {
	return m.repo.InsertUser(user)
}

func (m *UserModel) GetAllUser() ([]entity.User, error) {
	return m.repo.GetAllUser()
}

func (m *UserModel) GetUser(name string) (entity.User, error) {
	return m.repo.GetUser(name)
}
