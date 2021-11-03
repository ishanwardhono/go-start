package users

import (
	"app/database/repo"
	"app/entity"
	"app/errors"
	"context"
	"net/http"
	"strings"
)

type createUser struct {
	Req      entity.User
	RepoUser repo.UserRepo
}

func (m *createUser) Execute(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	return "", m.RepoUser.InsertUser(m.Req)
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
