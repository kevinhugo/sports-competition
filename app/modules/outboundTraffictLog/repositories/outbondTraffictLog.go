package repositories

import (
	"sports-competition/app/database"
	"sports-competition/app/helpers"
	outboundTraffictLogModel "sports-competition/app/modules/outboundTraffictLog/models"
	"time"

	"gorm.io/gorm"
)

type OutboundTraffictLogDBConnection struct {
	connection *gorm.DB
}

type OutboundTraffictLogRepository interface {
	SaveOutboundTraffictLog(createData map[string]interface{}) (*outboundTraffictLogModel.OutboundTraffictLog, error)
	UpdateOutboundTraffictLog(existingLog *outboundTraffictLogModel.OutboundTraffictLog, updateData map[string]interface{}) error
}

func NewOutboundTraffictLogRepository() OutboundTraffictLogRepository {
	return &OutboundTraffictLogDBConnection{
		connection: database.ConnectDB(),
	}
}

func (db *OutboundTraffictLogDBConnection) SaveOutboundTraffictLog(createData map[string]interface{}) (*outboundTraffictLogModel.OutboundTraffictLog, error) {
	var OutboundTraffictLog outboundTraffictLogModel.OutboundTraffictLog

	helpers.JsonToStruct(&createData, &OutboundTraffictLog)
	if OutboundTraffictLog.TrackingID == "" {
		OutboundTraffictLog.TrackingID = helpers.GetUUIDString()
	}
	var datetimeNow time.Time = helpers.GetDateTimeNow()
	OutboundTraffictLog.CreatedAt = datetimeNow

	err := db.connection.Model(&OutboundTraffictLog).Save(&OutboundTraffictLog).Error
	if err != nil {
		return nil, err
	}
	return &OutboundTraffictLog, nil
}

func (db *OutboundTraffictLogDBConnection) UpdateOutboundTraffictLog(existingLog *outboundTraffictLogModel.OutboundTraffictLog, updateData map[string]interface{}) error {
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
