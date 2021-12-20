package users

import (
	"app/database/repo/users"
	"app/entity"
	"app/module"
)

type Factory interface {
	Create(req entity.User) module.BaseModule
	Get(name string) module.BaseModule
	GetAll() module.BaseModule
}

type usersFactory struct {
	repoUser users.UserRepo
}

func NewUsersFactory(repoUser users.UserRepo) Factory {
	return &usersFactory{
		repoUser: repoUser,
	}
}

func (f *usersFactory) Create(req entity.User) module.BaseModule {
	return &createUser{
		Req:      req,
		RepoUser: f.repoUser,
	}
}

func (f *usersFactory) Get(name string) module.BaseModule {
	return &getUser{
		Req:      name,
		RepoUser: f.repoUser,
	}
}

func (f *usersFactory) GetAll() module.BaseModule {
	return &getUsers{
		RepoUser: f.repoUser,
	}
}
