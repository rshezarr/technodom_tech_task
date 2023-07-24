package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"redirect_api/internal/config"
	"redirect_api/internal/handler"
	"redirect_api/internal/repository"
	"redirect_api/internal/server"
	"redirect_api/internal/service"
	"redirect_api/pkg"
	"syscall"
)

func Start() {
	// connect to mongo DB
	db, err := pkg.ConnectDB(context.Background())
	if err != nil {
		log.Fatalf("error while connection db: %v", err)
	}

	// create collection
	if err := pkg.CreateCollections(context.Background(), db); err != nil {
		log.Fatalf("error while creating collection: %v", err)
	}

	// set config properties
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error while init configs: %v", err)
	}

	// get configs
	cfg := config.NewConfig()

	// set layers
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	// get server
	srv := server.NewServer(cfg, handler.InitRoutes())

	// graceful shutdown
	// creating channel for signals
	quit := make(chan os.Signal, 1)

	// accepting signal from user
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	// run server in goroutine
	go func() {
		log.Printf("app: Starting server at port %v -> http:// localhost%v\n", cfg.API.Port, cfg.API.Port)
		srv.Run()
	}()

	// "select" makes the server wait for any signal or error
	select {
	case sig := <-quit:
		log.Printf("app: signal accepted: %v\n", sig)
	case err := <-srv.ServerErrNotify():
		log.Printf("app: server closing: %v\n", err)
	}

	// shutting down the server when any error or signal occurs
	if err := srv.Shutdown(); err != nil {
		log.Printf("error while shutting down server: %s\n", err.Error())
	}
}
