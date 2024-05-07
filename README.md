##### WeatherApp 
##### Сервер на gin. В качестве БД - SQLite
```go
func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	router.GET("/:city", GetWeather)
	router.GET("/", Start)
	router.GET("/winterfell", EasterEgg)
	router.GET("/Winterfell", EasterEgg)
	
	return router
}
```



### TODO:
	1. Добавить функционал для бинарного файла ("city" в качестве аргумента или запрос от пользователя);
    2. Добавить больше информации о погоде (мб использовать другой сервис?);
    3. Добавить альтернативный сервис в случае недоступности основного;
    4. Покрытие тестами;
    5. Заменить БД на Redis.

### P.S.
try to find out the weather in Winterfell! :)
