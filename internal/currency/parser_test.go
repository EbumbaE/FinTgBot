package currency

import (
	"context"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	dbmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/storage"
)

func TestCancelParseCurrency(t *testing.T) {

	ctrl := gomock.NewController(t)
	ratesDB := dbmocks.NewMockRatesDB(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "allDoneWG", &sync.WaitGroup{})
	cancel()

	p, err := New(Config{Abbreviations: []string{}})
	assert.NoError(t, err)

	ratesDB.EXPECT().SetDefaultCurrency().Return(nil)

	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	err = p.ParseCurrencies(ctx, ratesDB)
	assert.NoError(t, err)
	ctx.Value("allDoneWG").(*sync.WaitGroup).Wait()
}
