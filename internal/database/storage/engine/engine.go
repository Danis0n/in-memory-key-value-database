package engine

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"hash/fnv"
)

type hashTable interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Del(key string)
}

type Engine struct {
	tables []hashTable
	logger *zap.Logger
}

func NewEngine(
	tableBuilder func() hashTable,
	logger *zap.Logger,
	tablesNumber int,
) (*Engine, error) {

	if logger == nil {
		return nil, errors.New("logger invalid")
	}

	if tablesNumber <= 0 {
		return nil, errors.New("tablesNumber invalid")
	}

	tables := make([]hashTable, tablesNumber)

	for i := 0; i < tablesNumber; i++ {
		if table := tableBuilder(); table != nil {
			tables[i] = table
		} else {
			return nil, errors.New("hash table table is invalid")
		}
	}

	return &Engine{
		tables: tables,
		logger: logger,
	}, nil
}

func (e *Engine) Set(ctx context.Context, key, value string) {
	tableID := e.calculateTableID(key)
	table := e.tables[tableID]
	table.Set(key, value)

	txID := ctx.Value("tx").(int64)
	e.logger.Debug("success set query", zap.Int64("tx", txID))
}

func (e *Engine) Get(ctx context.Context, key string) (string, bool) {
	tableID := e.calculateTableID(key)
	table := e.tables[tableID]
	get, found := table.Get(key)

	txID := ctx.Value("tx").(int64)
	e.logger.Debug("success get query", zap.Int64("tx", txID))

	return get, found
}

func (e *Engine) Del(ctx context.Context, key string) {
	tableID := e.calculateTableID(key)
	table := e.tables[tableID]
	table.Del(key)

	txID := ctx.Value("tx").(int64)
	e.logger.Debug("success del query", zap.Int64("tx", txID))
}

func (e *Engine) calculateTableID(key string) int {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(key))
	return int(hash.Sum32()) % len(e.tables)
}
