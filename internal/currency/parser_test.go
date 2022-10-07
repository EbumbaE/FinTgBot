package currency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCurrency(t *testing.T) {

	p, _ := New(Config{[]string{"USD", "CNY", "EUR"}})
	r, err := p.ParseCurrencies()
	assert.NoError(t, err)
	for valute := range r {
		fmt.Println(valute)
	}
}
