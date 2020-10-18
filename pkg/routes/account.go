package routes

import (
	"accounting/api/account"
	account2 "accounting/pkg/service/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

type balanceHandler struct {
	service account2.AccountService
}

func AddAccountHandler(e *gin.Engine, service account2.AccountService) {
	handler := &balanceHandler{
		service: service,
	}

	e.GET("/accounts", handler.GetAccount)
}

func (b *balanceHandler) GetAccount(ctx *gin.Context) {
	response := account.Account{
		Balance: b.service.Balance(),
	}

	ctx.JSON(http.StatusOK, response)
}
