package users

import (
	"app/database/repo/users"
	"app/entity"
	"app/errors"
	"context"
	"net/http"
	"strings"
)

type createUser struct {
	Req      entity.User
	RepoUser users.UserRepo
}

func (m *createUser) Execute(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	err = m.RepoUser.InsertUser(m.Req)
	if err != nil {
		return "insert failed", err
	}
	return "insert success", nil
}

func (m *createUser) Validate(ctx context.Context) error {
	var missingFields []string
	if m.Req.Name == "" {
		missingFields = append(missingFields, "name")
	}
	if m.Req.Email == "" {
		missingFields = append(missingFields, "email")
	}
	if len(missingFields) != 0 {
		return errors.New("missing fields: "+strings.Join(missingFields, ", "), http.StatusBadRequest)
	}
	return nil
}
