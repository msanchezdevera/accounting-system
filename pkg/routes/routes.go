package routes

import (
	account2 "accounting/pkg/service/account"
	"accounting/pkg/service/transaction"
	"github.com/gin-gonic/gin"
)

type HttpRoutes struct {
	transactionService transaction.TransactionService
	account            account2.AccountService
}

func NewHttpRoutes(transactionService transaction.TransactionService, accountService account2.AccountService) *HttpRoutes {
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
