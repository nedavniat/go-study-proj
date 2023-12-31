package main

import (
	"fmt"
	"github.com/vrischmann/envconfig"
	"example/go-study-proj/internal/router"
	c "example/go-study-proj/configs"
)

func main() {
	if err := envconfig.Init(&c.Conf); err != nil {
		fmt.Printf("err=%s\n", err)
	}

	router := router.NewRouter()
	router.Run("0.0.0.0:" + c.Conf.App.Port)
}
