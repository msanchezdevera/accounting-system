package routes

import (
	"accounting/api/balance"
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

	e.GET("/balance", handler.GetBalance)
}

func (b *balanceHandler) GetBalance(ctx *gin.Context) {
	response := balance.Balance{
		Balance: b.service.Balance(),
	}

	ctx.JSON(http.StatusOK, response)
}
