package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const URL = "https://api.open-meteo.com/v1/forecast?"
const PARAMS = "&current=temperature_2m,wind_speed_10m,wind_direction_10m&timezone=America/Sao_Paulo"

type WeatherAPIResponse struct {
	Current Current `json:"current"`
}

type Current struct {
	Time           string  `json:"time"`
	Interval       uint64  `json:"interval"`
	Temperature    float64 `json:"temperature_2m"`
	Wind_speed     float64 `json:"wind_speed_10m"`
	Wind_direction float64 `json:"wind_direction_10m"`
}

func GetWeather(c *gin.Context) {
	lat := c.Query("latitude")
	long := c.Query("longitude")

	latitude := fmt.Sprintf("latitude=%s", lat)
	longitude := fmt.Sprintf("longitude=%s", long)

	var data WeatherAPIResponse
	res, err := http.Get(URL + latitude + "&" + longitude + PARAMS)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "weather provider error"})
		return
	}
	
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
