package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

const parserTimeOut = time.Second * 10

type rateDB interface {
	AddRate(valute diary.Valute) error
	SetDefaultCurrency() error
}

type Parser struct {
	abbreviations []string
	urlCBR        string
}

type Response struct {
	Valute map[string]diary.Valute `json:"Valute"`
}

func New(config Config) (*Parser, error) {
	return &Parser{
		abbreviations: config.Abbreviations,
		urlCBR:        config.UrlCBR,
	}, nil
}

func (p *Parser) GetAbbreviations() []string {
	return p.abbreviations
}

func requestJsonCurrency(url string) ([]byte, error) {

	client := &http.Client{Timeout: parserTimeOut}

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
	jsonBytes, err := requestJsonCurrency(p.urlCBR)
	if err != nil {
		return err
	}

	valCurs := Response{}
	json.Unmarshal(jsonBytes, &valCurs)

	for _, abb := range p.abbreviations {
		if v, ok := valCurs.Valute[abb]; ok {
			if err = checkOnEmptyValues(v); err == nil {
				if err := storage.AddRate(v); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (p *Parser) ParseCurrencies(ctx context.Context, storage rateDB) error {

	storage.SetDefaultCurrency()

	go func() {

		timeTicker := time.NewTicker(time.Microsecond)

		defer timeTicker.Stop()

		for {
			select {
			case <-timeTicker.C:
				timeTicker = time.NewTicker(time.Hour * 24)

				if err := p.parseAndSendRates(storage); err != nil {
					log.Println("parse and send rates: ", err)
					return
				}

			case <-ctx.Done():
				defer ctx.Value("allDoneWG").(*sync.WaitGroup).Done()
				log.Println("parser is off")
				return
			}
		}
	}()

	return nil
}
