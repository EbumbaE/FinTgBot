package tgServer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Argument struct {
	lineArgs string
	amount   int64
}

type Testpair struct {
	values  Argument
	average []string
}

var parseTestsPair = []Testpair{
	{Argument{"a b 123 34.1 what", 5}, []string{"a", "b", "123", "34.1", "what"}},
	{Argument{"1", 1}, []string{"1"}},
	{Argument{"aasd 34,1 what", 3}, []string{"aasd", "34,1", "what"}},
}

func TestParseArguments(t *testing.T) {
	for _, pair := range parseTestsPair {
		result, err := parseArguments(pair.values.lineArgs, int(pair.values.amount))
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, result, pair.average, "expected: %p, but get %p", result, pair.average)
	}
}
