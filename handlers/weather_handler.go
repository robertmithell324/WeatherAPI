package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"weather-app/services"
	"weather-app/utils"
)


func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request from %s for %s", r.RemoteAddr, r.URL.Path)

	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		log.Printf("Error parsing latitude: %v", err)
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		log.Printf("Error parsing longitude: %v", err)
		return
	}

	apiKey := utils.GetAPIKey()
	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		log.Print("API key not set")
		return
	}

	weather, err := services.GetWeather(lat, lon, apiKey)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		log.Printf("Failed to fetch weather data: %v", err)
		return
	}

	weatherCondition := services.DetermineWeatherCondition(weather.Main.Temp)
	response := fmt.Sprintf("Weather: %s, Temperature: %.1f°C, Condition: %s", weather.Weather[0].Main, weather.Main.Temp, weatherCondition)

	log.Printf("Response sent for request from %s: %s", r.RemoteAddr, response)
	fmt.Fprintln(w, response)
}