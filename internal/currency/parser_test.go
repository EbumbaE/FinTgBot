package currency

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCurrency(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	abbreviations := []string{"USD", "CNY", "EUR"}
	p, err := New(Config{Abbreviations: abbreviations})
	assert.NoError(t, err)

	err = p.ParseCurrencies(ctx, nil)
	assert.NoError(t, err)

	cancel()
}
