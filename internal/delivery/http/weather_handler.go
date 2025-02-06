package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tomygp97/weather-notifier/internal/usecase"
)

type WeatherHandler struct {
	WeatherUsecase *usecase.WeatherUsecase
}

func NewWeatherHandler(weatherUsecase *usecase.WeatherUsecase) *WeatherHandler {
	return &WeatherHandler{
		WeatherUsecase: weatherUsecase,
	}
}

// GetWeather maneja la solicitud para obtener la previsión climática
func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityIDStr := vars["cityID"]
	cityID, err := strconv.Atoi(cityIDStr)
	if err != nil {
		http.Error(w, "ID de ciudad inválido", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	forecast, err := h.WeatherUsecase.GetWeather(ctx, cityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forecast)
}

// GetWaves maneja la solicitud para obtener la previsión de olas
func (h *WeatherHandler) GetWaves(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityIDStr := vars["cityID"]
	cityID, err := strconv.Atoi(cityIDStr)
	if err != nil {
		http.Error(w, "ID de ciudad inválido", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	waves, err := h.WeatherUsecase.GetWaves(ctx, cityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(waves)
}
