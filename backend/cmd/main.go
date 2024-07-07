package main

import (
	"app/backend/internal/config"
	"app/backend/internal/handler"
	"app/backend/internal/repository"
	"app/backend/internal/server"
	"app/backend/internal/service"
	"fmt"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg := config.MustLoad()

	db, err := repository.NewPostrgesDb(cfg.DataBase)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	rc, err := repository.NewRedisDb(cfg.RedisConfig)
	if err != nil {
		log.Fatalf("failed to connect to redis: %s", err)
	}

	_ = rc

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	sizeStorageSectors, err := repos.GetSizeStorageSectors()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	if sizeStorageSectors == 0 {
		if err = service.GetSectors(os.Getenv("SECTORS_PATH")); err != nil {
			log.Fatalf("failed to get sectors: %s", err)
		}
		fmt.Println("get sectors")
	}

	sizeStorageBillboards, err := repos.GetSizeStorageBillboards()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	if sizeStorageBillboards == 0 {
		if err = service.GetBillboards(os.Getenv("TRAIN_DATA_PATH")); err != nil {
			log.Fatalf("failed to get billboards: %s", err)
		}
		fmt.Println("get billboards")
	}

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.HTTPServer.Port, handler.InitRoutes()); err != nil {
			log.Fatalf("failed to run http server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("shutdown server")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("failed to shutdown: %s", err)
	}

	if err := db.Close(); err != nil {
		log.Fatalf("%s", err)
	}
}
