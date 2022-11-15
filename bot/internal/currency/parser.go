package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

const parserTimeOut = time.Second * 10
const parserTimer = time.Hour * 24

type rateDB interface {
	AddRate(valute diary.Valute) error
	SetDefaultCurrency() error
}

type Parser struct {
	abbreviations []string
	urlCBR        string
	client        *http.Client
}

type Response struct {
	Valute map[string]diary.Valute `json:"Valute"`
}

func New(config Config) (*Parser, error) {
	return &Parser{
		abbreviations: config.Abbreviations,
		urlCBR:        config.UrlCBR,
		client:        &http.Client{Timeout: parserTimeOut},
	}, nil
}

func (p *Parser) GetAbbreviations() []string {
	return p.abbreviations
}

func requestJsonCurrency(url string, client *http.Client) ([]byte, error) {

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.Status)
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func checkOnEmptyValues(v diary.Valute) error {
	if v.Abbreviation == "" || v.Name == "" || v.Value == 0 {
		return fmt.Errorf("Empty values")
	}
	return nil
}

func (p *Parser) parseAndSendRates(storage rateDB) error {
	jsonBytes, err := requestJsonCurrency(p.urlCBR, p.client)
	if err != nil {
		return err
	}

	valCurs := Response{}
	json.Unmarshal(jsonBytes, &valCurs)

	for _, abb := range p.abbreviations {
		if v, ok := valCurs.Valute[abb]; ok {
			err = checkOnEmptyValues(v)
			if err != nil {
				return err
			}
			err = storage.AddRate(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *Parser) ParseCurrencies(ctx context.Context, storage rateDB) error {

	storage.SetDefaultCurrency()

	logger.Info("currency parser begin")
	go func() {

		timeTicker := time.NewTicker(time.Microsecond)

		defer timeTicker.Stop()

		for {
			select {
			case <-timeTicker.C:
				timeTicker = time.NewTicker(parserTimer)

				if err := p.parseAndSendRates(storage); err != nil {
					logger.Error("parse and send rates: ", zap.Error(err))
					return
				}

			case <-ctx.Done():
				defer ctx.Value("allDoneWG").(*sync.WaitGroup).Done()
				logger.Info("currency parser end")
				return
			}
		}
	}()

	return nil
}
