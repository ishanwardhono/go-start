package main

import (
	"app/handler"
	"app/provider/database"
	"app/provider/database/repo"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	handler := handler.NewUserHandler(repo.NewUserRepo(database.GetDB()))
	handler.RegisterHandlers(router)
	defer database.GetDB().Close()

	fmt.Println("Starting localhost:8080 . . . ")
	log.Fatal(http.ListenAndServe(":8080", router))
}
