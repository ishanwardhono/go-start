package users

import (
	"app/database/repo"
	"app/entity"
	"context"
	"errors"
)

type createUser struct {
	Req      entity.User
	RepoUser repo.UserRepo
}

func (m *createUser) Execute(ctx context.Context) (interface{}, error) {
	m.Validate(ctx)
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	return "", m.RepoUser.InsertUser(m.Req)
}

func (m *createUser) Validate(ctx context.Context) error {
	if m.Req.Name == "" {
		return errors.New("name is mandatory")
	}
	if m.Req.Email == "" {
		return errors.New("email is mandatory")
	}
	return nil
}
