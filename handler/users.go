package handler

import (
	"app/entity"
	"app/log"
	"app/module/users"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	users users.Factory
}

func NewUserHandler(users users.Factory) HttpHandler {
	return &UserHandler{
		users: users,
	}
}

func (uh *UserHandler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/users", Handle(uh.allUsers)).Methods("GET")
	router.HandleFunc("/user/{name}", Handle(uh.getUser)).Methods("GET")
	router.HandleFunc("/user", Handle(uh.newUser)).Methods("POST")
}

func (uh *UserHandler) allUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return uh.users.GetAll().Execute(ctx)
}

func (uh *UserHandler) newUser(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var req entity.User

	err := ParseBody(r, &req)
	if err != nil {
		log.Error(ctx, "Failed parse request")
		return nil, err
	}

	return uh.users.Create(req).Execute(ctx)
}

func (uh *UserHandler) getUser(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return uh.users.Get(GetQueryParam(r, "name")).Execute(ctx)
}
