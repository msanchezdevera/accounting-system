package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	apiError "accounting/api/error"
	"accounting/pkg/errors"
	"accounting/pkg/log"
)

func NewErrorHandler(log log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()

		err := ctx.Errors.Last()
		if err == nil {
			return
		}

		apiError := newApiError(err.Err)

		log.Errorf("%s", err.Err)

		ctx.JSON(apiError.StatusCode, apiError)
	}
}

func newApiError(err error) apiError.Error {

	if customError, ok := err.(errors.Error); ok {
		var statusCode int

		switch customError.Type() {
		case errors.UserError:
			statusCode = http.StatusBadRequest
		case errors.NotFound:
			statusCode = http.StatusNotFound
		case errors.StatusUnsupportedMediaType:
			statusCode = http.StatusUnsupportedMediaType
		default:
			statusCode = http.StatusInternalServerError
		}

		return apiError.Error{
			Cause:      customError.Error(),
			Context:    customError.Context(),
			StatusCode: statusCode,
		}
	}

	return apiError.Error{
		Cause:      err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}
