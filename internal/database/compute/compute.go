package compute

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

type parser interface {
	ParseQuery(ctx context.Context, query string) ([]string, error)
}

type analyzer interface {
	AnalyzeQuery(ctx context.Context, tokens []string) (Query, error)
}

type Compute struct {
	analyzer analyzer
	parser   parser
	logger   *zap.Logger
}

func NewCompute(analyzer analyzer, parser parser, logger *zap.Logger) (*Compute, error) {
	if logger == nil {
		return nil, errors.New("invalid logger")
	}

	if parser == nil {
		return nil, errors.New("invalid parser")
	}

	if analyzer == nil {
		return nil, errors.New("invalid analyzer")
	}

	return &Compute{
		analyzer: analyzer,
		parser:   parser,
		logger:   logger,
	}, nil
}

func (c *Compute) HandleQuery(ctx context.Context, queryStr string) (Query, error) {
	tokens, err := c.parser.ParseQuery(ctx, queryStr)
	if err != nil {
		return Query{}, err
	}

	query, err := c.analyzer.AnalyzeQuery(ctx, tokens)
	if err != nil {
		return Query{}, err
	}

	return query, nil
}
