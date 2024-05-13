package repositories

import (
	"sports-competition/app/database"
	"sports-competition/app/helpers"
	inboundTraffictLogModel "sports-competition/app/modules/inboundTraffictLog/models"
	"time"

	"gorm.io/gorm"
)

type InboundTraffictLogDBConnection struct {
	connection *gorm.DB
}

type InboundTraffictLogRepository interface {
	SaveInboundTraffictLog(createData map[string]interface{}) (*inboundTraffictLogModel.InboundTraffictLog, error)
	UpdateInboundTraffictLog(existingLog *inboundTraffictLogModel.InboundTraffictLog, updateData map[string]interface{}) error
}

func NewInboundTraffictLogRepository() InboundTraffictLogRepository {
	return &InboundTraffictLogDBConnection{
		connection: database.ConnectDB(),
	}
}

func (db *InboundTraffictLogDBConnection) SaveInboundTraffictLog(createData map[string]interface{}) (*inboundTraffictLogModel.InboundTraffictLog, error) {
	var inboundTraffictLog inboundTraffictLogModel.InboundTraffictLog

	helpers.JsonToStruct(&createData, &inboundTraffictLog)
	if inboundTraffictLog.TrackingID == "" {
		inboundTraffictLog.TrackingID = helpers.GetUUIDString()
	}
	var datetimeNow time.Time = helpers.GetDateTimeNow()
	inboundTraffictLog.CreatedAt = datetimeNow

	err := db.connection.Model(&inboundTraffictLog).Save(&inboundTraffictLog).Error
	if err != nil {
		return nil, err
	}
	return &inboundTraffictLog, nil
}

func (db *InboundTraffictLogDBConnection) UpdateInboundTraffictLog(existingLog *inboundTraffictLogModel.InboundTraffictLog, updateData map[string]interface{}) error {
	_, isSet := updateData["updated_at"]
	if !isSet {
		updateData["updated_at"] = helpers.GetDateTimeNow()
	}
	helpers.JsonToStruct(&updateData, &existingLog)
	err := db.connection.Model(&existingLog).Save(&existingLog).Error
	if err != nil {
		return err
	}
	return nil
}
