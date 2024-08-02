package compute

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

var (
	errInvalidSymbol = errors.New("invalid symbol")
	errInvalidInput  = errors.New("invalid input: empty input")
)

type Parser struct {
	logger *zap.Logger
}

func NewParser(logger *zap.Logger) (*Parser, error) {
	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Parser{
		logger: logger,
	}, nil
}

func (p *Parser) ParseQuery(ctx context.Context, query string) ([]string, error) {
	if len(query) == 0 {
		return []string{}, errInvalidInput
	}

	machine := newStateMachine()

	tokens, err := machine.Parse(query)
	if err != nil {
		return nil, err
	}

	txID := ctx.Value("tx").(int64)

	p.logger.Debug(
		"tokens was parsed",
		zap.Int64("tx", txID),
		zap.Any("tokens", tokens),
	)

	return tokens, nil
}
