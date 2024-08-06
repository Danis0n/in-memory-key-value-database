package initialization

import (
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/console"
)

func CreateConsole(logger *zap.Logger) (*console.SimpleConsole, error) {
	return console.NewSimpleConsole(logger)
}
