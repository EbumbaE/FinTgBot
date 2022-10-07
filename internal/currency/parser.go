package currency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Parser struct {
	abbreviations []string
}

type Valute struct {
	Abbreviation string  `json:"CharCode"`
	Name         string  `json:"Name"`
	Value        float64 `json:"Value"`
}

type Responce struct {
	Valute map[string]Valute `json:"Valute"`
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

func checkOnEmptyValues(v Valute) error {
	if v.Abbreviation == "" || v.Name == "" || v.Value == 0 {
		return fmt.Errorf("Empty values")
	}
	return nil
}

func (p *Parser) ParseCurrencies() (chan Valute, error) {

	returnChan := make(chan Valute)
	go func() (err error) {
		for range time.Tick(time.Hour * 24) {
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
		}
		close(returnChan)
		return err
	}()

	return returnChan, nil
}
