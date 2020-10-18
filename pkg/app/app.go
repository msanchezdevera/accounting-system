package app

import (
	"accounting/pkg/config"
	"accounting/pkg/log"
	"accounting/pkg/routes"
	"accounting/pkg/server"
	"accounting/pkg/service"
	"accounting/pkg/storage"
)

type Application struct {
	log        log.Logger
	httpRoutes *routes.HttpRoutes
	server     *server.Server
	config     *config.Configuration
}

func NewApplication(config *config.Configuration, log log.Logger) *Application {

	transactionStorage := storage.NewTransaction()

	accountService := service.NewAccountService()

	transactionService := service.NewTransactionService(transactionStorage, accountService, log)

	httpRoutes := routes.NewHttpRoutes(transactionService, accountService)

	server := server.NewServer(config, log, httpRoutes)

	return &Application{
		log:        log,
		httpRoutes: httpRoutes,
		server:     server,
		config:     config,
	}
}

func (a *Application) Start() {
	a.server.SetShutdown(a.shutdown)
	a.server.Start()
}

func (a *Application) shutdown() {
}
