package main

import (
	"app/backend/internal/config"
	"app/backend/internal/models"
	"app/backend/internal/repository"
	"fmt"
	"log"
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

	id, err := repos.StorageBillboard.AddBillboard(models.Billboard{
		Lat:     "5243.52",
		Lon:     "14343.41",
		Azimuth: "90",
	})

	fmt.Println(err)
	fmt.Println(id)

	err = repos.DeleteBillboardById(5)

	fmt.Println(err)
}
