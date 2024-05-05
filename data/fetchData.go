package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testAgent/entities"
)

func FetchWeather(city string, ch chan<- entities.Weather) {
	defer close(ch)

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", entities.WeatherAPIKey, city)
	response, err := http.Get(url)
	entities.CommonError("Failed to fetch weather data\n", err)

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		entities.CommonError("Failed to fetch weather data\n"+response.Status, err)
	}

	var weatherData struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	}

	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		fmt.Printf("Failed to decode weather data: %v\n", err)
		return
	}

	weather := entities.Weather{
		City:        city,
		Temperature: weatherData.Current.TempC,
	}

	ch <- weather
}
