package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/esdrasbrz/primoflix/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	configJson, _ := json.Marshal(config)
	fmt.Println(string(configJson))
}
