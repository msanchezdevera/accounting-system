package middleware

import (
	"bytes"
	"io/ioutil"
	"accounting/pkg/log"

	"github.com/gin-gonic/gin"
)

func NewLogRequest(log log.Logger, skipPaths []string) gin.HandlerFunc {
	var skip map[string]struct{}
	if length := len(skipPaths); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range skipPaths {
			skip[path] = struct{}{}
		}
	}

	return func(ctx *gin.Context) {
		request := ctx.Request

		// Log only if the path isn't listed in skipped paths config
		if _, ok := skip[request.URL.Path]; !ok {
			var body []byte
			if request.Body != nil {
				body, _ = ioutil.ReadAll(request.Body)
			} else {
				body = make([]byte, 0)
			}

			request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			log.Debugf("[REQUEST] %s %s %s", request.Method, request.URL, string(body))
		}

		ctx.Next()
	}
}
