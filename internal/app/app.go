package app

import (
	"log"
	"os"
	"os/signal"
	"redirect_api/internal/config"
	"redirect_api/internal/handler"
	"redirect_api/internal/server"
	"syscall"
)

func Start() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error while init configs: %v", err)
	}

	cfg := config.NewConfig()

	handler := handler.NewHandler()

	server := server.NewServer(cfg, handler.InitRoutes())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		log.Printf("app: Starting server at port %v -> http://localhost%v\n", cfg.API.Port, cfg.API.Port)
		server.Run()
	}()

	select {
	case signal := <-quit:
		log.Printf("app: signal accepted: %v\n", signal)
	case err := <-server.ServerErrNotify():
		log.Printf("app: server closing: %v\n", err)
	}

	if err := server.Shutdown(); err != nil {
		log.Printf("%s\n", err.Error())
	}

}
