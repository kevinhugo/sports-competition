package routers

import (
	"fmt"
	"sports-competition/app/env"
	"sports-competition/app/logger"
	healthRouter "sports-competition/app/modules/health/routers"
	sportRouter "sports-competition/app/modules/sport/routers"
	userRouter "sports-competition/app/modules/user/routers"

	_ "sports-competition/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {
	logger.Info("Initializing routing start.")
	defer logger.Info("Initializing routing process has finished.")

	var api *gin.RouterGroup = r.Group(fmt.Sprintf("/%s", env.SERVICE_ROOT_PATH))
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	healthRouter.InitHealthRouter(api)
	userRouter.InitV1UserRouter(api)
	sportRouter.InitV1SportRouter(api)
}
