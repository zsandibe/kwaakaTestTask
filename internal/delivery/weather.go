package delivery

import (
	"fmt"
	"kwaaka-task/internal/domain"
	"kwaaka-task/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetWeather(c *gin.Context) {
	city := c.Query("city")

	if city == "" || !pkg.CityValidate(city) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city"})
		return
	}

	weather, err := h.service.GetWeatherByCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, weather)
}

func (h *handler) UpdateWeather(c *gin.Context) {
	var weatherRequest domain.Request

	if err := c.BindJSON(&weatherRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if weatherRequest.City == "" || !pkg.CityValidate(weatherRequest.City) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city"})
		return
	}
	if err := h.service.UpdateWeather(weatherRequest.City); err != nil {
		fmt.Println("ZDES")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Successfully updated")

}

func (h *handler) GetAllWeather(c *gin.Context) {
	weatherList, err := h.service.GetAllWeatherList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(weatherList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Weather list not found"})
		return
	}
	c.JSON(http.StatusOK, weatherList)
}
