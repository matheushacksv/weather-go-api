package main

import (
	"fmt"
	"log"
	"net/http"
	"weather/api/users"
	"weather/api/weather"
	"weather/config"
	"weather/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.Connect(cfg.DBString)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	router := gin.Default()

	// Health endpoint with ping in database
	router.GET("/health", func(c *gin.Context) {
		err := pool.Ping(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "down",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	})

	// User handler and endpoints
	userHandler := users.UserHandler{Pool: pool}
	router.POST("/users", userHandler.Create)
	router.GET("/users", userHandler.List)
	router.GET("/users/:ID", userHandler.Get)
	router.PUT("/users/:ID", userHandler.Update)
	router.DELETE("/users/:ID", userHandler.Delete)

	router.GET("/weather", weather.GetWeather)

	if err := router.Run(
		fmt.Sprintf(":%s", cfg.Port),
	); err != nil {
		log.Fatal("API Error")
	}
}
