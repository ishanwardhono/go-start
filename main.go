package main

import (
	"app/env"
	"app/handler"
	"app/log"
	"app/provider/config"
	"app/provider/database"
	"app/provider/database/repo"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	environment := flag.String("envi", "", "set environment")
	flag.Parse()
	env.SetEnv(*environment)

	cfg := config.GetConfig()
	log.Init(cfg.LogFile)

	router := mux.NewRouter().StrictSlash(true)
	handler := handler.NewUserHandler(repo.NewUserRepo(database.GetDB()))
	handler.RegisterHandlers(router)
	defer database.GetDB().Close()

	fmt.Println("Starting localhost:8080 . . . ")
	log.Fatal(http.ListenAndServe(":8080", router))
}
