package main

import (
	"context"
	"in-memory-key-value-database/internal/initialization"
	"os/signal"
	"syscall"
)

func run() error {
	return nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	initializer, err := initialization.NewInitializer()
	if err != nil {
		return
	}

	err = initializer.StartDatabase(ctx)
	if err != nil {
		return
	}

}
