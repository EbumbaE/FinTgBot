package currency

type Valute struct {
	Abbreviation string  `json:"CharCode"`
	Name         string  `json:"Name"`
	Value        float64 `json:"Value"`
}

func (v Valute) GetAbbreviation() string {
	return v.Abbreviation
}

func (v Valute) GetName() string {
	return v.Name
}

func (v Valute) GetValue() float64 {
	return v.Value
}
