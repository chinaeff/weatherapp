package main

import (
	"runtime/debug"
	"testAgent/data"
	"testAgent/repository"
)

func init() {
	debug.SetGCPercent(-1)
	repository.ConnectDB()
}

func main() {
	router := data.SetupRouter()
	router.Run(":8080")
}


