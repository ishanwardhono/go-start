package users

import (
	"app/database/repo"
	"app/errors"
	"context"
	"net/http"
)

type getUser struct {
	Req      string
	RepoUser repo.UserRepo
}

func (m *getUser) Execute(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	return m.RepoUser.GetUser(m.Req)
}

func (m *getUser) Validate(ctx context.Context) error {
	if m.Req == "" {
		return errors.New("name is not exist", http.StatusBadRequest)
	}
	return nil
}
