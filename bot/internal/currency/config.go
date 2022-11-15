package currency

type Config struct {
	Abbreviations []string `mapstructure:"abbreviations"`
	UrlCBR        string   `mapstructure:"urlCBR"`
}
