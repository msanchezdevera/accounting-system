package context

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"accounting/pkg/errors"
	"strings"
)

func DecodeBody(ctx *gin.Context, target interface{}) error {
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(target); err != nil {
		return errors.UserError.Wrapf(err, "Wrong JSON format: %v", ctx.Request.Body)
	}

	return nil
}

func CheckContentType(ctx *gin.Context) error {
	contentType := ctx.Request.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return errors.StatusUnsupportedMediaType.Newf("invalid Content-Type, expect `application/json`, got `%s`", contentType)
	}
	return nil
}
