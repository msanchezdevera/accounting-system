package routes

import (
	"accounting/api/account"
	"accounting/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type balanceHandler struct {
	account *model.Account
}

func AddAccountHandler(e *gin.Engine, account *model.Account) {
	handler := &balanceHandler{
		account: account,
	}

	e.GET("/accounts", handler.GetAccount)
}

func (b *balanceHandler) GetAccount(ctx *gin.Context) {
	response := account.Account{
		Balance: b.account.Balance(),
	}

	ctx.JSON(http.StatusOK, response)
}
