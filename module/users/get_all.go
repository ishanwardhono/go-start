package users

import (
	"app/provider/database/repo"
	"context"
)

type getUsers struct {
	RepoUser repo.UserRepo
}

func (m *getUsers) Execute(ctx context.Context) (interface{}, error) {
	m.Validate(ctx)
	return m.RepoUser.GetAllUser()
}

func (m *getUsers) Validate(ctx context.Context) error {
	return nil
}
