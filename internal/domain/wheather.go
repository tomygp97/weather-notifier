package domain

import "encoding/xml"

// WeatherForecast representa la estructura de la previsión climática
type WeatherForecast struct {
	XMLName   xml.Name   `xml:"cidade"`
	Name      string     `xml:"nome"`
	State     string     `xml:"uf"`
	UpdatedAt string     `xml:"atualizacao"`
	Forecasts []Forecast `xml:"previsao"`
}

type Forecast struct {
	Date    string  `xml:"dia"`
	Weather string  `xml:"tempo"`
	MaxTemp int     `xml:"maxima"`
	MinTemp int     `xml:"minima"`
	UVIndex float64 `xml:"iuv"`
}

// WavesForecast representa la estructura de la previsión de olas
type WavesForecast struct {
	XMLName xml.Name `xml:"cidade"`
	Waves   []Wave   `xml:"dados"`
}

type Wave struct {
	Date          string  `xml:"dia"`
	WaveHeight    float64 `xml:"altura"`
	WaveDirection string  `xml:"direcao"`
	WavePeriod    int     `xml:"periodo"`
	SeaCondition  string  `xml:"mar"`
}
