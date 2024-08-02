package compute

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	t.Parallel()

	_, err := NewAnalyzer(nil)
	require.Error(t, err)

	_, err = NewAnalyzer(zap.NewNop())
	require.NoError(t, err)
}

func TestAnalyze(t *testing.T) {
	t.Parallel()

	a, errA := NewAnalyzer(zap.NewNop())
	p, errP := NewParser(zap.NewNop())

	require.NoError(t, errA)
	require.NoError(t, errP)

	inputs := []TestInputs{
		{resultLen: 3, test: "SET 1 1", err: false},
		{resultLen: 2, test: "GET 1", err: false},
		{resultLen: 2, test: "DEL 2", err: false},
		{resultLen: 3, test: "DEL 2 2", err: true},
		{resultLen: 3, test: "DEL 2 2", err: true},
		{resultLen: 1, test: "SET    ", err: true},
	}

	ctx := context.WithValue(context.Background(), "tx", int64(555))

	for _, value := range inputs {

		tokens, err := p.ParseQuery(ctx, value.test)
		require.NoError(t, err)

		_, err = a.AnalyzeQuery(ctx, tokens)
		if value.err {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}

	}

}
