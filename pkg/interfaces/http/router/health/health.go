package health

import (
	"github.com/gin-gonic/gin"
	"github.com/thifnmi/go-book-api/pkg/interfaces/http/handlers"
)

func SetupRouter(router *gin.Engine, healthcheck handlers.HealthHandler) {
	api := router.Group("")
	{
		api.GET("/", healthcheck.Info)
		api.GET("/healthcheck", healthcheck.Heathcheck)
	}
}
