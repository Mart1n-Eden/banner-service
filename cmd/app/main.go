package main

import (
	"banner-service/internal/config"
	"banner-service/internal/handler"
	"banner-service/internal/repository"
	"banner-service/internal/server"
	"banner-service/internal/service"
	"context"
	_ "context"
	"os"
	"os/signal"
	_ "syscall"
	"time"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	cfg := config.NewConfig("configs/config.yaml")

	pg, err := repository.NewPostgres(cfg.DB)
	if err != nil {
		return
		// TODO: логирование
	}
	defer pg.Close()

	repo := repository.NewRepository(pg)
	srvc := service.NewService(repo)
	hand := handler.NewHandler(srvc)

	app := server.NewServer(hand.InitRoutes(), cfg.Server.Port)

	go func() {
		app.Run()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	app.Shutdown(ctx)
}
