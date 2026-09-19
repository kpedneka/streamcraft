package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kpedneka/streamcraft/internal/config"
	"github.com/kpedneka/streamcraft/internal/server"
	"github.com/kpedneka/streamcraft/internal/telemetry"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := telemetry.NewLogger(os.Stdout, "api")

	dependencies := server.Dependencies{
		// Initialize any dependencies required by the server here.
	}

	router := server.NewRouter(dependencies)
	addr := fmt.Sprintf(":%s", config.HTTPPort)
	logger.Info("api starting", "addr", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Error("api failed to start", "error", err)
		os.Exit(1)
	}
}
