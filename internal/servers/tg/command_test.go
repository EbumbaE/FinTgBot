package tgServer

import (
	"testing"
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

func Comparation(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func Test_ParseArguments(t *testing.T) {
	for _, pair := range parseTestsPair {
		result, err := parseArguments(pair.values.lineArgs, int(pair.values.amount))
		if err != nil {
			t.Error(err)
		}
		if !Comparation(result, pair.average) {
			t.Errorf("expected: %p, but get %p", result, pair.average)
		}
	}
}
