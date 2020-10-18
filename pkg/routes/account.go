package routes

import (
	"accounting/api/account"
	"accounting/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type balanceHandler struct {
	service service.AccountService
}

func AddAccountHandler(e *gin.Engine, service service.AccountService) {
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
