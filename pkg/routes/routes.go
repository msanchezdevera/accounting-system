package routes

import (
	"accounting/pkg/service"
	"github.com/gin-gonic/gin"
)

type HttpRoutes struct {
	transactionService service.TransactionService
	account            service.AccountService
}

func NewHttpRoutes(transactionService service.TransactionService, accountService service.AccountService) *HttpRoutes {
	return &HttpRoutes{
		transactionService: transactionService,
		account:            accountService,
	}
}

func (r *HttpRoutes) AddHttpRoutes(e *gin.Engine) {
	addHealthCheckRoutes(e)
	r.addApplicationRoutes(e)
}

func (r *HttpRoutes) addApplicationRoutes(e *gin.Engine) {
	AddTransactionHandler(e, r.transactionService)
	AddAccountHandler(e, r.account)
}
