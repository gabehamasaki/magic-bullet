package main

import (
	"log"
	"os"

	api "magic-bullet/backend/internal/api"
	"magic-bullet/backend/internal/config"
	"magic-bullet/backend/internal/database"
)

func main() {

	config := config.NewConfig()
	if err := config.Load(); err != nil {
		log.Fatalf("failed do load config: %v", err)
	}

	database, err := database.NewPGSQLConnector(config)
	if err != nil {
		log.Fatalf("failed do create database connector: %v", err)
	}

	app := api.NewApp(config, database)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "openapi":
			if err := api.WriteOpenAPI(os.Stdout, app.API); err != nil {
				log.Fatalf("failed to generate openapi spec: %v", err)
			}
			return
		case "serve":
			// Explicit serve command for parity with the default behavior.
		default:
			log.Fatalf("unknown command %q\nusage: go run ./cmd/api [serve|openapi]", os.Args[1])
		}
	}

	addr := api.HTTPAddr()
	log.Printf("magic-bullet backend listening on %s", addr)
	if err := app.Engine.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
