package routers

import (
	"sports-competition/app/middlewares"
	sportHandlers "sports-competition/app/modules/sport/handlers"

	"github.com/gin-gonic/gin"
)

func InitV1SportRouter(router *gin.RouterGroup) {
	var sportHandler *sportHandlers.SportHandler = sportHandlers.NewSportHandler()

	var v1Router *gin.RouterGroup = router.Group("/v1")
	{
		var sportRouter *gin.RouterGroup = v1Router.Group("/sport")
		{
			sportRouter.POST("/begin-competition", middlewares.AuthMiddleware(), middlewares.RequestLogger("Begin Competition"), sportHandler.BeginCompetition)
		}
	}
}
