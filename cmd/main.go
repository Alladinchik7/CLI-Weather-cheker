package main

import (
	"CLIproject/api/wether"
	"CLIproject/inernal/city"
	display "CLIproject/inernal/displayWeather"
	"CLIproject/pkg/system"
	"fmt"
	"os"
)

func main() {
	// Читаем API-ключ из .env
	apiKey := system.LogApiKey()
	if  apiKey == "" {
		fmt.Println("Error: Please set the OPENWEATHER_API_KEY environment variable.")
		os.Exit(1)
	}

	// Проверяем, передано ли название города в аргументах командной строки
	city, err := city.ChekCity()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Получаем данные о погоде
	weather, err := weather.GetetWeatherData(city, apiKey)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}

	display.DisplayWeatherData(weather)
}