package main

import (
	"app/env"
	"app/handler"
	"app/log"
	"app/provider/config"
	"app/provider/database"
	"app/provider/database/repo"
	"flag"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	environment := flag.String("env", "", "set environment")
	flag.Parse()
	env.SetEnv(*environment)

	cfg := config.GetConfig()
	log.Init(cfg.LogFile)

	router := mux.NewRouter().StrictSlash(true)
	handler := handler.NewUserHandler(repo.NewUserRepo(database.GetDB()))
	handler.RegisterHandlers(router)
	defer database.GetDB().Close()

	log.Info(nil, "Server running port "+cfg.AppPort+" on "+env.GetEnv()+" . . . ")
	log.Fatal(nil, http.ListenAndServe(":"+cfg.AppPort, router))
}
