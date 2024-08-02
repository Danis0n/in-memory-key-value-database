package initialization

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/database"
	"in-memory-key-value-database/internal/database/compute"
	"in-memory-key-value-database/internal/database/storage"
	"in-memory-key-value-database/internal/database/storage/engine"
	"in-memory-key-value-database/internal/run"
)

type runner interface {
	HandleQueries(ctx context.Context) error
}

type Initializer struct {
	engine engine.Engine
	logger *zap.Logger
}

func createEngine(logger *zap.Logger) (*engine.Engine, error) {
	return engine.NewEngine(engine.HashTableBuilder, logger, 10)
}

func (i *Initializer) createStorageLayer(logger *zap.Logger) (*storage.Storage, error) {
	return storage.NewStorage(logger, &i.engine)
}

func (i *Initializer) createComputeLayer(logger *zap.Logger) (*compute.Compute, error) {
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

// here may be some manipulation with configuration
func (i *Initializer) createRunner(databaseLayer database.Database) (runner, error) {
	return run.NewConsoleRunner(databaseLayer)
}

func NewInitializer() (*Initializer, error) {

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	en, err := createEngine(logger)
	if err != nil {
		return nil, err
	}

	return &Initializer{
		engine: *en,
		logger: logger,
	}, nil
}

func (i *Initializer) StartDatabase(ctx context.Context) error {
	storageLayer, err := i.createStorageLayer(i.logger)
	if err != nil {
		return err
	}

	computeLayer, err := i.createComputeLayer(i.logger)
	if err != nil {
		return err
	}

	databaseLayer, err := database.NewDatabase(i.logger, storageLayer, computeLayer)
	if err != nil {
		return err
	}

	runner, consErr := i.createRunner(*databaseLayer)

	if consErr != nil {
		return consErr
	}

	for {
		err := runner.HandleQueries(ctx)
		if err != nil {
			fmt.Print(err)
		}
	}
}
