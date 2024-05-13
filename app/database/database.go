package database

import (
	"fmt"
	"sports-competition/app/env"
	"time"

	inboundTraffictLogModel "sports-competition/app/modules/inboundTraffictLog/models"
	outboundTraffictLogModel "sports-competition/app/modules/outboundTraffictLog/models"
	sportModel "sports-competition/app/modules/sport/models"
	userModel "sports-competition/app/modules/user/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(additionalConfig ...any) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", env.DB_HOST, env.DB_USER, env.DB_PASS, env.DB_NAME, env.DB_PORT)
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect database")
	}

	var pingDB bool = true

	var additionalConfigLength int = len(additionalConfig)
	if additionalConfigLength > 0 {
		pingDB = additionalConfig[0].(bool)
	}

	if pingDB {
		go PingDB(db)
	}

	return db
}

func MigrateDatabase(db *gorm.DB) {
	userModel.MigrateUserStatusEnum(db)

	db.AutoMigrate(&inboundTraffictLogModel.InboundTraffictLog{})
	db.AutoMigrate(&outboundTraffictLogModel.OutboundTraffictLog{})
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&sportModel.Sport{})
}

/*
Ping DB Connection regularly to avoid connection lost because of idle.
*/
func PingDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error while trying to get database connection, see logs below.")
		fmt.Println(err)
	} else {
		for {
			time.Sleep(time.Duration(env.DB_PING_DELAY_IN_SECONDS) * time.Second)
			err = sqlDB.Ping()
			if err != nil {
				fmt.Println("Error while trying to ping database connection, see logs below.")
				fmt.Println(err)
			}
		}
	}
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
