package handler

import (
	"app/entity"
	"app/log"
	"app/module"
	"app/provider/database/repo"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	user module.UserModel
}

func NewUserHandler(repo repo.UserRepo) UserHandler {
	return UserHandler{
		user: module.NewUserModel(repo),
	}
}

func (uh *UserHandler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/users", Handle(uh.allUsers)).Methods("GET")
	router.HandleFunc("/user/{name}", Handle(uh.getUser)).Methods("GET")
	router.HandleFunc("/user", Handle(uh.newUser)).Methods("POST")
}

func (uh *UserHandler) allUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return uh.user.GetAllUser()
}

func (uh *UserHandler) newUser(tx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var req entity.User

	err := ParseBody(r, &req)
	if err != nil {
		log.Error("Failed parse request")
		return nil, err
	}

	return uh.user.InsertUser(req)
}

func (uh *UserHandler) getUser(tx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return uh.user.GetUser(GetQueryParam(r, "name"))
}
