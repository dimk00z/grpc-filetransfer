package main

import (
	"log"

	config "github.com/dimk00z/grpc-filetransfer/config/server"
	"github.com/dimk00z/grpc-filetransfer/internal/server/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	log.Println("config: ", cfg)
	// Run
	app.Run(cfg)
}
