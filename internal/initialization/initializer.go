package initialization

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/configuration"
	"in-memory-key-value-database/internal/database"
	"in-memory-key-value-database/internal/database/compute"
	"in-memory-key-value-database/internal/database/storage"
	"in-memory-key-value-database/internal/network"
)

type Runner interface {
	Start(ctx context.Context, handler network.Handler) error
}

type Initializer struct {
	dbEngine storage.Engine
	logger   *zap.Logger
	runner   Runner
}

func NewInitializer(cfg *configuration.Configuration) (*Initializer, error) {
	if cfg == nil {
		return nil, errors.New("configuration invalid")
	}

	logger, err := CreateLogger(cfg.LoggingConfiguration)
	if err != nil {
		return nil, err
	}

	dbEngine, err := CreateEngine(cfg.EngineConfiguration, logger)
	if err != nil {
		return nil, err
	}

	runner, err := CreateRunner(logger, cfg.RunnerConfiguration, cfg.NetworkConfiguration)
	if err != nil {
		return nil, err
	}

	return &Initializer{
		dbEngine: dbEngine,
		logger:   logger,
		runner:   runner,
	}, nil
}

func (i *Initializer) StartDatabase(ctx context.Context) error {
	storageLayer, err := createStorageLayer(i.logger, i.dbEngine)
	if err != nil {
		return err
	}

	computeLayer, err := createComputeLayer(i.logger)
	if err != nil {
		return err
	}

	db, err := database.NewDatabase(i.logger, storageLayer, computeLayer)
	if err != nil {
		return err
	}

	err = i.runner.Start(ctx, func(ctx context.Context, queryStr []byte) []byte {
		response := db.HandleQuery(ctx, string(queryStr))
		return []byte(response)
	})
	if err != nil {
		return err
	}

	return nil
}

func createStorageLayer(logger *zap.Logger, dbEngine storage.Engine) (*storage.Storage, error) {
	return storage.NewStorage(logger, dbEngine)
}

func createComputeLayer(logger *zap.Logger) (*compute.Compute, error) {
	parser, err := compute.NewParser(logger)
	if err != nil {
		return nil, err
	}

	analyzer, err := compute.NewAnalyzer(logger)
	if err != nil {
		return nil, err
	}

	return compute.NewCompute(analyzer, parser, logger)
}
