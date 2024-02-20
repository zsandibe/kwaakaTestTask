package delivery

import "github.com/gin-gonic/gin"

func (h *handler) Routes() *gin.Engine {
	router := gin.Default()
	router.PUT("/weather", h.UpdateWeather)
	router.GET("/weather", h.GetWeather)
	return router
}
