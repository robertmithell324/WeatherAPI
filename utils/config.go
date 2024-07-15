package utils

import (
	"log"
	"os"
)


func GetAPIKey() string {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Print("API key not found in environment variables")
	}
	return apiKey
}