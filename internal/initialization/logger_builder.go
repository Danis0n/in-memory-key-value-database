package initialization

import (
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/configuration"
)

// CreateLogger create logger
func CreateLogger(cfg *configuration.LoggingConfiguration) (*zap.Logger, error) {
	if cfg == nil {
		return zap.NewDevelopment()
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
