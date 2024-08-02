package compute

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestParserInitFalse(t *testing.T) {
	_, err := NewParser(nil)

	require.Error(t, err)
}

func TestParser(t *testing.T) {
	t.Parallel()

	p, err := NewParser(zap.NewNop())

	require.NoError(t, err)

	inputs := []TestInputs{
		{resultLen: 3, test: "SET 1 1"},
		{resultLen: 2, test: "GET 1"},
		{resultLen: 2, test: "DEL 2"},
		{resultLen: 3, test: "DEL 2 2"},
		{resultLen: 3, test: "DEL 2 2"},
		{resultLen: 1, test: "SET    "},
	}

	for _, value := range inputs {
		ctx := context.WithValue(context.Background(), "tx", int64(555))

		parse, err := p.ParseQuery(ctx, value.test)

		require.NoError(t, err)
		require.Equal(t, len(parse), value.resultLen)

	}

}
