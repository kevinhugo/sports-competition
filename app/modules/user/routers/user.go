package routers

import (
	"sports-competition/app/middlewares"
	userHandlers "sports-competition/app/modules/user/handlers"

	"github.com/gin-gonic/gin"
)

func InitV1UserRouter(router *gin.RouterGroup) {
	var userHandler *userHandlers.UserHandler = userHandlers.NewUserHandler()

	var v1Router *gin.RouterGroup = router.Group("/v1")
	{
		var userRouter *gin.RouterGroup = v1Router.Group("/user")
		{
			userRouter.POST("/login", userHandler.Login)
			userRouter.GET("/identity/:first_name", middlewares.AuthMiddleware(), middlewares.RequestLogger("Get User Identity"), userHandler.GetUserIdentity)
		}
	}
}
