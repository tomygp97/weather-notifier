package usecase

import (
	"encoding/json"
	"fmt"
	"time"

	"context"

	"github.com/redis/go-redis/v9"
	"github.com/tomygp97/weather-notifier/internal/domain"
	"github.com/tomygp97/weather-notifier/internal/infrastructure/repository"
)

type WeatherUsecase struct {
	weatherRepo *repository.WeatherRepository
	redisClient *redis.Client
}

func NewWeatherUsecase(weatherRepo *repository.WeatherRepository, redisClient *redis.Client) *WeatherUsecase {
	return &WeatherUsecase{
		weatherRepo: weatherRepo,
		redisClient: redisClient,
	}
}

// GetWeather obtiene la previsión climática, usando Redis como caché
func (uc *WeatherUsecase) GetWeather(ctx context.Context, cityID int) (*domain.WeatherForecast, error) {
	cacheKey := fmt.Sprintf("weather:%d", cityID)
	cachedData, err := uc.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var forecast domain.WeatherForecast
		if err := json.Unmarshal([]byte(cachedData), &forecast); err == nil {
			return &forecast, nil
		}
	} else if err != redis.Nil {
		// Si ocurre un error distinto a que la clave no exista, lo manejamos
		return nil, fmt.Errorf("error al obtener datos de Redis: %w", err)
	}

	forecastResponse, err := uc.weatherRepo.FetchWeather(ctx, cityID)
	if err != nil {
		return nil, err
	}

	// Convert forecastResponse to domain.WeatherForecast
	weatherForecast := &domain.WeatherForecast{
		XMLName:   forecastResponse.XMLName,
		Name:      forecastResponse.Name,
		State:     forecastResponse.State,
		UpdatedAt: forecastResponse.UpdatedAt,
		Forecasts: forecastResponse.Forecasts,
	}

	// Guardar en caché por 30 minutos
	data, _ := json.Marshal(weatherForecast)
	uc.redisClient.Set(ctx, cacheKey, data, 30*time.Minute)

	return weatherForecast, nil
}

// GetWaves obtiene la previsión de olas, usando Redis como caché
func (uc *WeatherUsecase) GetWaves(ctx context.Context, cityID int) (*domain.WavesForecast, error) {
	cacheKey := fmt.Sprintf("waves:%d", cityID)
	cachedData, err := uc.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var waves domain.WavesForecast
		if err := json.Unmarshal([]byte(cachedData), &waves); err == nil {
			return &waves, nil
		}
	} else if err != redis.Nil {
		// Si ocurre un error distinto a que la clave no exista, lo manejamos
		return nil, fmt.Errorf("error al obtener datos de Redis: %w", err)
	}

	dayOffset := 0 // Asignar el valor adecuado
	wavesResponse, err := uc.weatherRepo.FetchWaves(ctx, cityID, dayOffset)
	if err != nil {
		return nil, err
	}

	// Convert wavesResponse to domain.WavesForecast
	wavesForecast := &domain.WavesForecast{
		XMLName: wavesResponse.XMLName,
		Waves:   wavesResponse.Waves,
	}

	// Guardar en caché por 30 minutos
	data, _ := json.Marshal(wavesForecast)
	uc.redisClient.Set(ctx, cacheKey, data, 30*time.Minute)

	return wavesForecast, nil
}
