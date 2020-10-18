package routes

import (
	apiError "accounting/api/error"
	"accounting/api/transaction"
	"accounting/pkg/errors"
	"accounting/pkg/log"
	"accounting/pkg/model"
	"accounting/pkg/routes/mocks"
	"accounting/test_fixture"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//go:generate mockgen -destination=mocks/mock_TransactionService.go -package=mocks -source=../service/transaction.go

type transactionMocks struct {
	ctrl    *gomock.Controller
	service *mocks.MockTransactionService
}

func (builder *transactionMocks) build() *gin.Engine {
	router := test_fixture.SetupRouter(log.NewConfigless())
	AddTransactionHandler(router, builder.service)
	return router
}

func transactionSetUp(t *testing.T) (*gin.Engine, *transactionMocks) {
	ctrl := gomock.NewController(t)
	mocks := &transactionMocks{
		ctrl:    ctrl,
		service: mocks.NewMockTransactionService(ctrl),
	}
	return mocks.build(), mocks
}

var effectiveDate, _ = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2020")

var createTransactionAPI = transaction.CreateTransaction{
	Amount: 100,
	Type:   "credit",
}

var transactionAPI = transaction.Transaction{
	ID:            "test-id",
	Amount:        100,
	Type:          "credit",
	EffectiveDate: effectiveDate,
}

var transactionModel = model.Transaction{
	ID:              "test-id",
	TransactionType: model.Credit,
	Amount:          100,
	EffectiveDate:   effectiveDate,
}

func TestTransactionService_Create(t *testing.T) {
	t.Run("Create - transaction ok", func(t *testing.T) { transaction_create_success(t) })
	t.Run("Create - transaction create error", func(t *testing.T) { transactionService_create_serverError(t) })
	t.Run("Create - transaction create content type error", func(t *testing.T) { transactionService_create_contentTypeError(t) })
}

func TestTransactionService_Get(t *testing.T) {
	t.Run("Get - transaction ok", func(t *testing.T) { transaction_create_success(t) })
	t.Run("Get - transaction create error", func(t *testing.T) { transactionService_create_serverError(t) })
}

func transaction_create_success(t *testing.T) {
	router, mocks := transactionSetUp(t)
	defer mocks.ctrl.Finish()

	mocks.service.EXPECT().
		Create(&createTransactionAPI).
		Return(&transactionModel, nil).
		Times(1)

	bodyStr, _ := json.Marshal(createTransactionAPI)
	request, response := test_fixture.NewRequest("POST", "/transactions", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	actualResponseBody := transaction.Transaction{}
	json.NewDecoder(response.Body).Decode(&actualResponseBody)

	assert.Equal(t, 200, response.Code)

	test_fixture.Diff(t, transactionAPI, actualResponseBody)
}

func transactionService_create_serverError(t *testing.T) {
	router, mocks := transactionSetUp(t)
	defer mocks.ctrl.Finish()

	mocks.service.EXPECT().
		Create(&createTransactionAPI).
		Return(nil, errors.UserError.New("some strange error")).
		Times(1)

	bodyStr, _ := json.Marshal(createTransactionAPI)
	request, response := test_fixture.NewRequest("POST", "/transactions", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	actualResponseBody := apiError.Error{}
	json.NewDecoder(response.Body).Decode(&actualResponseBody)

	expectedResponseBody := apiError.Error{
		Cause:      "some strange error",
		StatusCode: 400,
	}

	assert.Equal(t, 400, response.Code)
	test_fixture.Diff(t, expectedResponseBody, actualResponseBody)
}

func transactionService_create_contentTypeError(t *testing.T) {
	router, mocks := transactionSetUp(t)
	defer mocks.ctrl.Finish()

	bodyStr, _ := json.Marshal(createTransactionAPI)
	request, response := test_fixture.NewRequest("POST", "/transactions", bytes.NewBuffer(bodyStr))
	request.Header.Set("Content-Type", "wrong")

	router.ServeHTTP(response, request)

	actualResponseBody := apiError.Error{}
	json.NewDecoder(response.Body).Decode(&actualResponseBody)

	expectedResponseBody := apiError.Error{
		Cause:      "invalid Content-Type, expect `application/json`, got `wrong`",
		StatusCode: 415,
	}

	assert.Equal(t, 415, response.Code)
	test_fixture.Diff(t, expectedResponseBody, actualResponseBody)
}
