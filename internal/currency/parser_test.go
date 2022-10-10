package currency

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnClosedChannelParseCurrency(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	p, err := New(Config{[]string{"USD", "CNY", "EUR"}})
	assert.NoError(t, err)

	rateCh, err := p.ParseCurrencies(ctx)
	assert.NoError(t, err)

	cancel()
	for valute := range rateCh {
		err := fmt.Errorf("channel is open, we read: %v", valute)
		assert.NoError(t, err)
	}
}

func TestParseCurrency(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	abbreviations := []string{"USD", "CNY", "EUR"}
	p, err := New(Config{abbreviations})
	assert.NoError(t, err)

	r, err := p.ParseCurrencies(ctx)
	assert.NoError(t, err)

	mapAbb := make(map[string]struct{})
	countAbb := 0
	var emp struct{}
	for valute := range r {
		mapAbb[valute.Abbreviation] = emp
		countAbb++
		if countAbb == len(abbreviations) {
			cancel()
		}
	}

	for _, x := range abbreviations {
		_, ok := mapAbb[x]
		assert.Equal(t, ok, true)
	}
}
