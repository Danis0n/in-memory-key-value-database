package storage

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

type engine interface {
	Set(ctx context.Context, key, value string)
	Get(ctx context.Context, key string) (string, bool)
	Del(ctx context.Context, key string)
}

type Storage struct {
	logger *zap.Logger
	engine engine
}

func NewStorage(logger *zap.Logger, engine engine) (*Storage, error) {
	if engine == nil {
		return nil, errors.New("engine invalid")
	}

	if logger == nil {
		return nil, errors.New("logger invalid")
	}

	return &Storage{
		logger: logger,
		engine: engine,
	}, nil
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	s.engine.Set(ctx, key, value)
	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	value, _ := s.engine.Get(ctx, key)
	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	s.engine.Del(ctx, key)
	return nil
}
