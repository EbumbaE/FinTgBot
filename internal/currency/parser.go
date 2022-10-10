package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type Parser struct {
	abbreviations []string
}

type Responce struct {
	Valute map[string]diary.Valute `json:"Valute"`
}

func New(config Config) (*Parser, error) {
	return &Parser{
		abbreviations: config.Abbreviations,
	}, nil
}

func (p *Parser) GetAbbreviations() []string {
	return p.abbreviations
}

func requestJsonCurrency() ([]byte, error) {

	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://www.cbr-xml-daily.ru/daily_json.js"

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

func (p *Parser) ParseCurrencies(ctx context.Context) (chan diary.Valute, error) {

	returnChan := make(chan diary.Valute)
	go func() {

		timeTicker := time.NewTicker(time.Microsecond)

		defer timeTicker.Stop()
		defer close(returnChan)

		for {
			select {
			case <-timeTicker.C:
				timeTicker = time.NewTicker(time.Second * 3)

				jsonBytes, err := requestJsonCurrency()
				if err != nil {
					break
				}

				valCurs := Responce{}
				json.Unmarshal(jsonBytes, &valCurs)

				for _, abb := range p.abbreviations {
					if v, ok := valCurs.Valute[abb]; ok {
						if err = checkOnEmptyValues(v); err == nil {
							returnChan <- v
						}
					}
				}
			case <-ctx.Done():
				log.Println("parser is off")
				return
			}
		}
	}()

	return returnChan, nil
}
