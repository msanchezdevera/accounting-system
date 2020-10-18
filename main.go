package main

import (
	"accounting/pkg/app"
	"accounting/pkg/config"
	"accounting/pkg/log"
)

func main() {

	config := config.ParseConfig()

	logger := log.NewLogger(config)

	logger.Infof("Launching accounting system")

	application := app.NewApplication(config, logger)

	application.Start()

}
