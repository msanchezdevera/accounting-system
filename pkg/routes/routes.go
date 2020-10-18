package routes

import (
	"accounting/pkg/model"
	"accounting/pkg/service"
	"github.com/gin-gonic/gin"
)

type HttpRoutes struct {
	transactionService service.TransactionService
	account            *model.Account
}

func NewHttpRoutes(transactionService service.TransactionService, account *model.Account) *HttpRoutes {
	return &HttpRoutes{
		transactionService: transactionService,
		account:            account,
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
