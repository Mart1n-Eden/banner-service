package main

import (
	"banner-service/internal/config"
	"banner-service/internal/handler"
	"banner-service/internal/repository"
	"banner-service/internal/server"
	"banner-service/internal/service"
	_ "context"
	"os"
	"os/signal"
	_ "syscall"
)

func main() {
	cfg := config.NewConfig("configs/config.yaml")
	
	pg, err := repository.NewPostgres(cfg.DB)
	if err != nil {
		return
		// TODO: логирование
	}
	defer pg.Close()

	repository := repository.NewRepository(pg)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	app := server.NewServer(handler.InitRoutes(), cfg.Server.Port)

	go func() {
		app.Run()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
}
