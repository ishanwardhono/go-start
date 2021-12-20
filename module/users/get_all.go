package users

import (
	"app/database/repo/users"
	"context"
)

type getUsers struct {
	RepoUser users.UserRepo
}

func (m *getUsers) Execute(ctx context.Context) (interface{}, error) {
	m.Validate(ctx)
	return m.RepoUser.GetAllUser()
}

func (m *getUsers) Validate(ctx context.Context) error {
	return nil
}
