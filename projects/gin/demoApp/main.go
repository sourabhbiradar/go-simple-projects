package main

import (
	"demoApp/config"
	"demoApp/logger"
	"demoApp/router"
)

func main() {
	// demo app using logrus

	config.Init()
	config.Appconfig = config.GetConfig()

	logger.Init()
	logger.InfoLn("logger initialized successfully")
	logger.InfoLn("started router initialization")

	router.Init()

	logger.InfoLn("router initialized successfully")
}
