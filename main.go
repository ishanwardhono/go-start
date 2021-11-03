package main

import (
	"app/config"
	"app/env"
	"app/log"
	"app/provider"
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	environment := flag.String("env", "", "set environment")
	flag.Parse()
	env.SetEnv(*environment)

	cfg := config.GetConfig()
	log.Init(cfg.LogLevel, cfg.LogFile)

	handlers := provider.GetHandlers()
	router := mux.NewRouter().StrictSlash(true)
	for _, handler := range handlers {
		handler.RegisterHandlers(router)
	}

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Info(context.Background(), "[MAIN] Server running port "+cfg.AppPort+" on "+env.GetEnv()+" . . . ")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(context.Background(), "[MAIN] Server Error !!! err:", err)
		}
	}()

	<-sign
	log.Info(context.Background(), "[MAIN] Server Stopping . . .")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.MaxGraceStop)*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(ctx, "[MAIN] Server Shutdown Failed !!! err:", err)
	}
	log.Info(ctx, "[MAIN] Server Stopped Gracefully")
}
