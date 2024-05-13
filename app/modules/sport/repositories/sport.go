package repositories

import (
	"sports-competition/app/database"
	"sports-competition/app/helpers"
	sportModel "sports-competition/app/modules/sport/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sportDBConnection struct {
	connection *gorm.DB
}

type SportRepository interface {
	GetByUserID(userID string) (*sportModel.Sport, error)
	AddSport(userID string, proficiency uint) (*sportModel.Sport, error)
	UpdateSport(existingSport *sportModel.Sport, updateData map[string]any) error
}

func NewSportRepository() SportRepository {
	return &sportDBConnection{
		connection: database.ConnectDB(),
	}
}

func (db *sportDBConnection) GetByUserID(userID string) (*sportModel.Sport, error) {
	var sport sportModel.Sport

	err := db.connection.Model(&sport).
		Where("users.user_id = ?", userID).
		Joins("JOIN users ON sports.user_id = users.id").
		Find(&sport).
		Error
	if err != nil {
		return nil, err
	}

	return &sport, nil
}

func (db *sportDBConnection) AddSport(userID string, proficiency uint) (*sportModel.Sport, error) {
	var sport sportModel.Sport

	var createData map[string]any = map[string]any{}
	var datetimeNow time.Time = helpers.GetDateTimeNow()
	createData["created_at"] = datetimeNow
	createData["updated_at"] = datetimeNow
	createData["user_id"] = clause.Expr{SQL: "(SELECT id FROM users WHERE user_id = ?)", Vars: []any{userID}}
	createData["proficiency"] = proficiency

	err := db.connection.Model(&sport).Create(&createData).Error
	if err != nil {
		return nil, err
	}
	helpers.JsonToStruct(&createData, &sport)
	return &sport, nil
}

func (db *sportDBConnection) UpdateSport(existingSport *sportModel.Sport, updateData map[string]any) error {
	if _, isSet := updateData["updated_at"]; !isSet {
		updateData["updated_at"] = helpers.GetDateTimeNow()
	}

	helpers.JsonToStruct(&updateData, &existingSport)

	err := db.connection.Model(&existingSport).Save(&existingSport).Error
	if err != nil {
		return err
	}
	return nil
}
