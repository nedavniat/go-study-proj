package main

import (
	"example/go-study-proj/router"
)

func main() {
	router := router.NewRouter()
	router.Run("localhost:8080")
}
