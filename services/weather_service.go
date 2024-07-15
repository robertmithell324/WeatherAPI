package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-app/utils"
)

type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}


func GetWeather(lat, lon float64, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&APPID=%s", lat, lon, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}

	return &weatherResponse, nil
}

func DetermineWeatherCondition(temp float64) string {
	switch {
	case temp >= 30:
		return "hot"
	case temp <= 10:
		return "cold"
	default:
		return "moderate"
	}
}