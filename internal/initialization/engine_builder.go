package initialization

import (
	"errors"
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/configuration"
	"in-memory-key-value-database/internal/database/storage"
	"in-memory-key-value-database/internal/database/storage/engine/in_memory"
)

const (
	InMemoryEngine = "in_memory"
)

var engineTypes = map[string]struct{}{
	InMemoryEngine: {},
}

func CreateEngine(cfg *configuration.EngineConfiguration, logger *zap.Logger) (storage.Engine, error) {
	if cfg == nil {
		return in_memory.NewEngine(in_memory.HashTableBuilder, logger, 10)
	}

	if cfg.Type != "" {
		if _, ok := engineTypes[cfg.Type]; !ok {
			return nil, errors.New("unknown dbEngine type: " + cfg.Type)
		}
	}

	return in_memory.NewEngine(in_memory.HashTableBuilder, logger, 10)
}
