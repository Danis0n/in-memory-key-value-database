package initialization

import (
	"errors"
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/configuration"
)

const (
	NetworkRunner = "network"
	SimpleRunner  = "simple"
)

func CreateRunner(
	logger *zap.Logger,
	cfg *configuration.RunnerConfiguration,
	networkCfg *configuration.NetworkConfiguration,
) (Runner, error) {
	if logger == nil {
		return nil, errors.New("logger invalid")
	}

	if cfg == nil {
		return CreateConsole(logger)
	}

	if cfg.Type == SimpleRunner {
		return CreateConsole(logger)
	} else if cfg.Type == NetworkRunner {
		return CreateNetwork(networkCfg, logger)
	} else {
		return nil, errors.New("runner type invalid")
	}

}
