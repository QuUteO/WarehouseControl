package main

import (
	"WarehouseControl/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.New("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
