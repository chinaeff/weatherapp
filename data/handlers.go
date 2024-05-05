package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"testAgent/entities"
	"testAgent/repository"
	"time"
)

func GetWeather(c *gin.Context) {
	city := c.Param("city")

	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan entities.Weather, 1)

	go FetchWeather(city, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case weather := <-ch:
		c.JSON(http.StatusOK, fmt.Sprintf("Temperature in %v: %v degrees Celsius", weather.City, weather.Temperature))
		fmt.Printf("Temperature in %v: %v degrees Celsius\n", weather.City, weather.Temperature)

		var count int
		row := repository.DB.QueryRow("SELECT COUNT(*) FROM weather WHERE city = ?", weather.City)
		if err := row.Scan(&count); err != nil {
			log.Printf("Failed to check data: %v\n", err)
			return
		}

		if count == 0 {
			_, err := repository.DB.Exec("INSERT INTO weather(city, temperature) VALUES (?, ?)", weather.City, weather.Temperature)
			entities.CommonError("Failed to insert weather data\n", err)
		}
	case <-time.After(5 * time.Second):
		c.String(http.StatusInternalServerError, "Timeout error")
	}
}

func Start(c *gin.Context) {
	c.JSON(http.StatusOK, "Add '/city' to address, where 'city' is a city name")
	fmt.Println("Add '/city' to address in your web browser, where 'city' is a city name")
}

func EasterEgg(c *gin.Context) {
	c.File("repository/bin/winterfell.jpg")
	fmt.Println("See on your web browser: " + "http://localhost:8080/winterfell")
}
