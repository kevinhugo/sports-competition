package main

import (
	"fmt"
	"sports-competition/app/database"
	"sports-competition/app/env"
	"sports-competition/app/logger"
	"sports-competition/app/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	logger.Info(fmt.Sprintf("========== Initalization %s service start ==========", env.SERVICE_ROOT_PATH))
	var db *gorm.DB = database.ConnectDB()
	database.MigrateDatabase(db)
	database.DisconnectDB(db)
	r := gin.Default()
	routers.InitRouter(r)
	logger.Info(fmt.Sprintf("========== Initialization %s service has finished ==========", env.SERVICE_ROOT_PATH))
	r.Run(fmt.Sprintf("%s:%s", env.SERVICE_HOST, env.SERVICE_PORT))
}
