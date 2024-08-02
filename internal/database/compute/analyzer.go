package compute

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

var (
	SetArgumentsNumber = 2
	GetArgumentsNumber = 1
	DelArgumentsNumber = 1
)

var queryArgumentsNumber = map[int]int{
	SetCommandID: SetArgumentsNumber,
	GetCommandID: GetArgumentsNumber,
	DelCommandID: DelArgumentsNumber,
}

var (
	errInvalidCommand   = errors.New("invalid command")
	errInvalidArguments = errors.New("invalid arguments")
)

type Analyzer struct {
	logger *zap.Logger
}

func NewAnalyzer(logger *zap.Logger) (*Analyzer, error) {
	if logger == nil {
		return nil, errors.New("logger invalid")
	}

	return &Analyzer{
		logger: logger,
	}, nil
}

func (a *Analyzer) AnalyzeQuery(ctx context.Context, tokens []string) (Query, error) {
	if len(tokens) == 0 {
		txID := ctx.Value("tx").(int64)
		a.logger.Debug("invalid query", zap.Int64("tx", txID))
		return Query{}, errInvalidCommand
	}

	command := tokens[0]
	commandID := CommandFromName(command)

	if commandID == UnknownCommandID {
		txID := ctx.Value("tx").(int64)
		a.logger.Debug(
			"unknown command",
			zap.Int64("tx", txID),
			zap.Any("command", command),
		)
		return Query{}, errInvalidCommand
	}

	query := NewQuery(commandID, tokens[1:])
	argumentsNumber := queryArgumentsNumber[commandID]

	if len(query.Arguments) != argumentsNumber {
		txID := ctx.Value("tx").(int64)
		a.logger.Debug(
			"invalid arguments for query",
			zap.Int64("tx", txID),
			zap.Any("arguments", query.Arguments),
		)
		return Query{}, errInvalidArguments
	}

	txID := ctx.Value("tx").(int64)
	a.logger.Debug(
		"query analyzed",
		zap.Int64("tx", txID),
		zap.Any("Command", tokens),
	)

	return query, nil
}
