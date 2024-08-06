package main

import (
	"context"
	"in-memory-key-value-database/internal/configuration"
	"in-memory-key-value-database/internal/initialization"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ConfigFileName = os.Getenv("CONFIG_FILE_NAME")
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := &configuration.Configuration{}

	if ConfigFileName != "" {
		var err error
		cfg, err = configuration.Load(ConfigFileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	initializer, err := initialization.NewInitializer(cfg)
	if err != nil {
		return
	}

	err = initializer.StartDatabase(ctx)
	if err != nil {
		return
	}

}
