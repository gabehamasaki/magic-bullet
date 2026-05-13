package main

import (
	"log"
	"os"

	mbapi "magic-bullet/backend/internal/api"
)

func main() {
	app := mbapi.NewApp()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "openapi":
			if err := mbapi.WriteOpenAPI(os.Stdout, app.API); err != nil {
				log.Fatalf("failed to generate openapi spec: %v", err)
			}
			return
		case "serve":
			// Explicit serve command for parity with the default behavior.
		default:
			log.Fatalf("unknown command %q\nusage: go run ./cmd/api [serve|openapi]", os.Args[1])
		}
	}

	addr := mbapi.HTTPAddr()
	log.Printf("magic-bullet backend listening on %s", addr)
	if err := app.Engine.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
