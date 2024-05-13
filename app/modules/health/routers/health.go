package routers

import (
	healthHandlers "sports-competition/app/modules/health/handlers"

	"github.com/gin-gonic/gin"
)

func InitHealthRouter(router *gin.RouterGroup) {
	var healthHandler *healthHandlers.HealthHandler = healthHandlers.NewHealthHandler()

	router.GET("/health", healthHandler.Health)
}
