package routes

import (
	"accounting/api/account"
	"accounting/pkg/log"
	"accounting/pkg/model"
	"accounting/test_fixture"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type accountMocks struct {
	ctrl    *gomock.Controller
	account *model.Account
}

func (builder *accountMocks) build() *gin.Engine {
	router := test_fixture.SetupRouter(log.NewConfigless())
	AddAccountHandler(router, builder.account)
	return router
}

func accountSetUp(t *testing.T) (*gin.Engine, *accountMocks) {
	ctrl := gomock.NewController(t)
	mocks := &accountMocks{
		ctrl:    ctrl,
		account: model.NewAccount(),
	}

	return mocks.build(), mocks
}

func TestAccountService_Get(t *testing.T) {
	t.Run("Account - get ok", func(t *testing.T) { account_get_success(t) })
}

func account_get_success(t *testing.T) {
	router, mocks := accountSetUp(t)
	defer mocks.ctrl.Finish()

	mocks.account.UpdateBalance(100)

	request, response := test_fixture.NewRequest("GET", "/accounts", nil)

	router.ServeHTTP(response, request)

	actualResponseBody := account.Account{}
	json.NewDecoder(response.Body).Decode(&actualResponseBody)

	expectedResponse := account.Account{Balance: 100}

	assert.Equal(t, 200, response.Code)

	test_fixture.Diff(t, expectedResponse, actualResponseBody)
}
