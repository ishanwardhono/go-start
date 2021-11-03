package main

import (
	"app/env"
	"app/log"
	"app/provider"
	"app/provider/config"
	"context"
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

	handlers := provider.GetHandlers()
	router := mux.NewRouter().StrictSlash(true)
	for _, handler := range handlers {
		handler.RegisterHandlers(router)
	}

	log.Info(context.Background(), "Server running port "+cfg.AppPort+" on "+env.GetEnv()+" . . . ")
	log.Fatal(context.Background(), http.ListenAndServe(":"+cfg.AppPort, router))
}
