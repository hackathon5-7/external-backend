package main

import (
	"app/backend/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
