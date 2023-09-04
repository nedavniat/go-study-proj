package main

import (
	"example/go-study-proj/router"
)

func main() {
	router := router.NewRouter()
	router.Run("0.0.0.0:8080")
}
