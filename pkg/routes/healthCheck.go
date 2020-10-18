package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HealthCheckPath = "/health-check"
)

func addHealthCheckRoutes(e *gin.Engine) {
	e.GET(HealthCheckPath, healthCheck)
}

func healthCheck(c *gin.Context) {
	c.Header("no-gzip", "health-check-disabled-gzip")
	c.Header("cache-control", "no-cache")
	c.String(http.StatusOK, "OK")
}
