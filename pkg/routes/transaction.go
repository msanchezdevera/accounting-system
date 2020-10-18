package routes

import (
	"accounting/api/transaction"
	"accounting/pkg/errors"
	"accounting/pkg/model"
	"accounting/pkg/server/context"
	transaction2 "accounting/pkg/service/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transactionHandler struct {
	transactionService transaction2.TransactionService
}

func AddTransactionHandler(e *gin.Engine, service transaction2.TransactionService) {
	handler := &transactionHandler{
		transactionService: service,
	}

	e.POST("/transactions", handler.Create)
	e.GET("/transactions", handler.GetAll)
	e.GET("/transactions/:id", handler.Get)
}

// Create commits a new transaction to the account
func (t *transactionHandler) Create(ctx *gin.Context) {
	if err := context.CheckContentType(ctx); err != nil {
		ctx.Error(err)
		return
	}

	var createTransaction transaction.CreateTransaction
	var transactionApi transaction.Transaction

	if err := context.DecodeBody(ctx, &createTransaction); err != nil {
		ctx.Error(err)
		return
	}

	if err := t.validate(createTransaction); err != nil {
		ctx.Error(err)
		return
	}

	if response, err := t.transactionService.Create(&createTransaction); err != nil {
		ctx.Error(err)
		return
	} else {
		transactionApi = t.transactionToApi(response)
	}

	ctx.JSON(http.StatusOK, transactionApi)
}

// Get finds transaction by ID
func (t *transactionHandler) Get(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	if modelTransaction := t.transactionService.Get(transactionId); modelTransaction != nil {
		response := t.transactionToApi(modelTransaction)
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.Error(errors.NotFound.Newf("transaction %s not found", transactionId))
	}
}

// GetAll fetches transactions history
func (t *transactionHandler) GetAll(ctx *gin.Context) {
	var response []transaction.Transaction

	for _, transact := range t.transactionService.GetAll() {
		response = append(response, t.transactionToApi(transact))
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *transactionHandler) transactionToApi(model *model.Transaction) transaction.Transaction {
	return transaction.Transaction{
		ID:            model.ID,
		Amount:        model.Amount,
		Type:          string(model.TransactionType),
		EffectiveDate: model.EffectiveDate,
	}
}

func (t *transactionHandler) validate(createTransaction transaction.CreateTransaction) errors.Error {
	if err := model.TransactionType(createTransaction.Type).IsValid(); err != nil {
		return err
	}

	if createTransaction.Amount < 0 {
		return errors.UserError.New("transaction amount should be greater than zero")
	}

	return nil
}
