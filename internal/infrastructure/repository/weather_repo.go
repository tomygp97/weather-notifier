package repository

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/tomygp97/weather-notifier/internal/domain"
	"golang.org/x/net/html/charset"
)

const (
	baseWeatherURL = "http://servicos.cptec.inpe.br/XML/cidade/%d/previsao.xml"
	baseWavesURL   = "http://servicos.cptec.inpe.br/XML/cidade/%d/dia/%d/ondas.xml"
)

type WeatherRepository struct{}

func NewWeatherRepository() *WeatherRepository {
	return &WeatherRepository{}
}

type WeatherResponse struct {
	XMLName   xml.Name          `xml:"cidade"`
	Name      string            `xml:"nome"`
	State     string            `xml:"uf"`
	UpdatedAt string            `xml:"atualizacao"`
	Forecasts []domain.Forecast `xml:"previsao"`
}

type WavesResponse struct {
	XMLName xml.Name      `xml:"cidade"`
	Name    string        `xml:"nome"`
	State   string        `xml:"uf"`
	Waves   []domain.Wave `xml:"dados>ondas"`
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// FetchWeather obtiene la previsión climática de una ciudad por ID
func (r *WeatherRepository) FetchWeather(ctx context.Context, cityID int) (*WeatherResponse, error) {
	url := fmt.Sprintf(baseWeatherURL, cityID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud HTTP: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al obtener datos de clima: código de estado %d", resp.StatusCode)
	}

	var weather WeatherResponse
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&weather); err != nil {
		return nil, fmt.Errorf("error al decodificar el XML de clima: %w", err)
	}

	if weather.Name == "" || len(weather.Forecasts) == 0 {
		return nil, errors.New("datos de clima incompletos en la respuesta")
	}

	return &weather, nil
}

// FetchWaves obtiene la previsión de olas si la ciudad es costera
func (r *WeatherRepository) FetchWaves(ctx context.Context, cityID int, dayOffset int) (*WavesResponse, error) {
	url := fmt.Sprintf(baseWavesURL, cityID, dayOffset)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud HTTP: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al obtener datos de olas: código de estado %d", resp.StatusCode)
	}

	var wavesResponse WavesResponse
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&wavesResponse); err != nil {
		return nil, fmt.Errorf("error al decodificar el XML de olas: %w", err)
	}

	if len(wavesResponse.Waves) == 0 {
		return nil, errors.New("no hay datos de olas disponibles para esta ciudad y día")
	}

	return &wavesResponse, nil
}
