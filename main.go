package main

import (
	"fmt"
	"log"

	"github.com/esdrasbrz/primoflix/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*config.Radarr, *config.Sonarr)
}
