package main

import (
	"github.com/juliopragana/goestudo/config"
	"github.com/juliopragana/goestudo/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Config
	err := config.Init()
	//Verifiy erro initialize config
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
