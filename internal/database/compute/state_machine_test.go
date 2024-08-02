package compute

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type TestInputs struct {
	resultLen int
	test      string
	err       bool
}

func readyTestValues() []TestInputs {

	return []TestInputs{
		{resultLen: 3, test: "SET 1 1"},
		{resultLen: 2, test: "GET 1"},
		{resultLen: 2, test: "DEL 2"},
		{resultLen: 3, test: "DEL 2 2"},
		{resultLen: 3, test: "DEL 2 2"},
		{resultLen: 1, test: "SET    "},
	}
}

func TestStateMachineParse(t *testing.T) {

	values := readyTestValues()

	for _, val := range values {
		sm := newStateMachine()
		parse, err := sm.Parse(val.test)

		require.NoError(t, err)
		require.Equal(t, len(parse), val.resultLen)
	}

}
