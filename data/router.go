package data

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/:city", GetWeather)
	router.GET("/", Start)
	router.GET("/winterfell", EasterEgg)
	router.GET("/Winterfell", EasterEgg)

	return router
}
