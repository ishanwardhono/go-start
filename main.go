package main

import (
	"app/core/config"
	"app/core/log"
	"app/env"
	"app/provider"
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	//set environment
	environment := flag.String("env", "", "set environment")
	flag.Parse()
	env.SetEnv(*environment)

	//set config
	cfg := config.GetConfig()
	log.Init(cfg.LogLevel, cfg.LogFile)

	//register http server
	handlers := provider.GetHttpHandlers()
	router := mux.NewRouter().StrictSlash(true)
	for _, handler := range handlers {
		handler.RegisterHandlers(router)
	}
	httpServer := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	//register grpc server
	lis, _ := net.Listen("tcp", ":8081")
	grpcServer := provider.GetGrpcServers()

	//run http server
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Info(context.Background(), "[MAIN] HTTP Server running port "+cfg.AppPort+" on "+env.GetEnv()+" . . . ")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(context.Background(), "[MAIN] HTTP Server Error !!! err:", err)
		}
	}()

	go func() {
		log.Info(context.Background(), "[MAIN] gRPC Server running port "+cfg.AppPort+" on "+env.GetEnv()+" . . . ")
		if err := grpcServer.Serve(lis); err != nil && err != http.ErrServerClosed {
			log.Fatal(context.Background(), "[MAIN] gRPC HTTP Server Error !!! err:", err)
		}
	}()

	//gracfully stopping
	<-sign
	log.Info(context.Background(), "[MAIN] Server Stopping . . .")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.MaxGraceStop)*time.Second)
	defer cancel()

	//shutdown server
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(ctx, "[MAIN] Server Shutdown Failed !!! err:", err)
	}
	log.Info(ctx, "[MAIN] Server Stopped Gracefully")
}
