package run

import (
	"bufio"
	"context"
	"fmt"
	"in-memory-key-value-database/internal/database"
	"os"
)

type ConsoleRunner struct {
	databaseLayer database.Database
	run           func() error
}

func (cr *ConsoleRunner) HandleQueries(ctx context.Context) error {

	reader := bufio.NewReader(os.Stdin)
	queryStr, _ := reader.ReadString('\n')

	result := cr.databaseLayer.HandleQuery(ctx, queryStr)
	fmt.Print(result)

	return nil
}

func NewConsoleRunner(databaseLayer database.Database) (*ConsoleRunner, error) {

	return &ConsoleRunner{
		databaseLayer: databaseLayer,
		run:           func() error { return nil },
	}, nil
}
