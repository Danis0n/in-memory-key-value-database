package console

import (
	"bufio"
	"context"
	"fmt"
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/network"
	"os"
)

type SimpleConsole struct {
	logger *zap.Logger
}

func NewSimpleConsole(logger *zap.Logger) (*SimpleConsole, error) {
	return &SimpleConsole{
		logger: logger,
	}, nil
}

func (cr *SimpleConsole) Start(ctx context.Context, handler network.Handler) error {

	for {
		err := cr.handleQueries(ctx, handler)
		if err != nil {
			cr.logger.Error("handle queries", zap.Error(err))
			return nil
		}

	}

}

func (cr *SimpleConsole) handleQueries(ctx context.Context, handler network.Handler) error {
	reader := bufio.NewReader(os.Stdin)
	queryStr, _ := reader.ReadString('\n')

	result := handler(ctx, []byte(queryStr))
	fmt.Print(string(result), "\n")

	return nil
}
