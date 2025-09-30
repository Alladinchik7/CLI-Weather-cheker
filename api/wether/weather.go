package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Address string `json:"address"`
	CurrentConditions struct {
		Temp float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"currentConditions"`
	Alerts []struct {
		Description string `json:"description"`
	} `json:"alerts"`
}

func GetetWeatherData(city, apiKey string) (*WeatherData, error){
	// Формируем URL API
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/?key=%s&&unitGroup=metric", city, apiKey)

	// Отправляем HTTP-запрос
	resp, err := http.Get(url)
	if err != nil {
		return  nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Читаем что запрос прошёл
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: received status code %d %s", resp.StatusCode, string(bodyBytes))
	}

	// Читаем тело ответа
	bodyBites, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Декодируем JSON в структуру WeatherData
	var  weather WeatherData
	err = json.Unmarshal(bodyBites, &weather)
	if err != nil {
		return nil, fmt.Errorf("failed to decod JSON: %v", err)
	}

	return &weather, nil
}