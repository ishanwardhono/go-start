package main

import (
	"fmt"
	"log"
	"net/http"
	"sm-secret/database"
	"sm-secret/database/repo"
	"sm-secret/handler"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	handler := handler.NewUserHandler(repo.NewUserRepo(database.NewDB()))
	handler.RegisterHandlers(router)
	defer database.NewDB().Close()

	fmt.Println("Starting localhost:8080 . . . ")
	log.Fatal(http.ListenAndServe(":8080", router))
}
