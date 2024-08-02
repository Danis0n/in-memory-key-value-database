package database

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/database/compute"
)

type storageLayer interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type computeLayer interface {
	HandleQuery(context.Context, string) (compute.Query, error)
}

type Database struct {
	logger      *zap.Logger
	idGenerator *IDGenerator
	storage     storageLayer
	compute     computeLayer
}

func NewDatabase(
	logger *zap.Logger,
	storage storageLayer,
	compute computeLayer,
) (*Database, error) {

	if compute == nil {
		return nil, errors.New("compute is invalid")
	}

	if storage == nil {
		return nil, errors.New("storage is invalid")
	}

	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Database{
		logger:      logger,
		idGenerator: NewIDGenerator(),
		storage:     storage,
		compute:     compute,
	}, nil
}

func (d *Database) HandleQuery(ctx context.Context, queryStr string) string {
	txID := d.idGenerator.Generate()
	ctx = context.WithValue(ctx, "tx", txID)

	d.logger.Debug(
		"handling query",
		zap.Int64("tx", txID),
		zap.String("query", queryStr),
	)

	query, err := d.compute.HandleQuery(ctx, queryStr)
	if err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	switch query.Command {
	case compute.SetCommandID:
		return d.Set(ctx, query)
	case compute.GetCommandID:
		return d.Get(ctx, query)
	case compute.DelCommandID:
		return d.Del(ctx, query)
	}

	return "[error] configuration failure"
}

func (d *Database) Set(ctx context.Context, query compute.Query) string {
	arguments := query.Arguments
	err := d.storage.Set(ctx, arguments[0], arguments[1])
	if err != nil {
		return "[error] bad set"
	}

	return "[ok]"
}

func (d *Database) Get(ctx context.Context, query compute.Query) string {
	arguments := query.Arguments
	value, err := d.storage.Get(ctx, arguments[0])
	if err != nil {
		return "[error] bad get"
	}

	return fmt.Sprintf("[ok] %s", value)
}

func (d *Database) Del(ctx context.Context, query compute.Query) string {
	arguments := query.Arguments
	err := d.storage.Del(ctx, arguments[0])
	if err != nil {
		return "[error] bad del"
	}

	return "[ok]"
}
