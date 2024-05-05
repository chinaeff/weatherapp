*Необходимо написать API-сервис, который будет выдавать информацию из внешних источников(курс валют, погода, список фильмов и тд). В приложении необходимо использовать go-рутины и любое хранилище данных(psql, redis, mongo, etc)


##### Сервер на gin. В качестве БД - SQLite
```go
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/:city", GetWeather)
	router.GET("/", Start)

	return router
}
```



### TODO:
	1. Добавить функционал для бинарного файла ("city" в качестве аргумента или запрос от пользователя);
    2. Добавить больше информации о погоде (мб использовать другой сервис?);
    3. Покрытие тестами;
    4. Заменить БД на Redis.

### P.S.
try to find out the weather in Winterfell! :)