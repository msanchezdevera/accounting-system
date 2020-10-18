package test_fixture

import (
	"accounting/pkg/errors"
	"accounting/pkg/log"
	"accounting/pkg/server/middleware"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

var SomeError = errors.New("some strange error")

var BadJSON = "{:,}"

func SetupRouter(log log.Logger) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(middleware.NewLogRequest(log, []string{}))

	r.Use(middleware.NewErrorHandler(log))

	return r
}

func Diff(t *testing.T, x, y interface{}, opts ...cmp.Option) {
	if diff := cmp.Diff(x, y, opts...); diff != "" {
		t.Errorf(diff)
	}
}

func NewRequest(method, url string, body io.Reader) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(method, url, body)
	request.Header.Add("Content-Type", "application/json")
	return request, httptest.NewRecorder()
}
