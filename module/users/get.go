package users

import (
	"app/database/repo/users"
	"app/errors"
	"context"
	"net/http"
)

type getUser struct {
	Req      string
	RepoUser users.UserRepo
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
		return errors.New("missing field name", http.StatusBadRequest)
	}
	return nil
}
