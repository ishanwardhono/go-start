package handler

import (
	"app/database/repo"
	"app/entity"
	"app/module"
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/users", uh.allUsers).Methods("GET")
	router.HandleFunc("/user/{name}", uh.getUser).Methods("GET")
	router.HandleFunc("/user", uh.newUser).Methods("POST")
}

func (uh *UserHandler) allUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.user.GetAllUser()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(users)
}

func (uh *UserHandler) newUser(w http.ResponseWriter, r *http.Request) {
	var req entity.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Fprintf(w, "can't parse request")
		return
	}

	err = uh.user.InsertUser(req)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Success")
}

func (uh *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	queryParam := mux.Vars(r)
	user, err := uh.user.GetUser(queryParam["name"])
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
}
