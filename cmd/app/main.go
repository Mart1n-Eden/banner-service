package main

import (
	"banner-service/internal/cache"
	"banner-service/internal/config"
	"banner-service/internal/handler"
	"banner-service/internal/repository"
	"banner-service/internal/server"
	"banner-service/internal/service"
	"context"
	_ "context"
	"fmt"
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
		fmt.Println("error conection db")
		fmt.Println(err.Error())
		return
		// TODO: логирование
	}
	defer pg.Close()

	red, err := cache.NewRedis()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("error conection red")
		return
		// TODO: логирование
	}
	defer red.Close()

	repo := repository.NewRepository(pg)
	cache := cache.NewCache(red)
	srvc := service.NewService(repo, cache)
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

//import (
//	"context"
//	"fmt"
//	"github.com/go-redis/redis/v8"
//)
//
//var ctx = context.Background()
//
//func main() {
//	ExampleClient()
//}
//
//func ExampleClient() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}
