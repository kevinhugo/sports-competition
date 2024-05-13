package repositories

import (
	"sports-competition/app/database"
	"sports-competition/app/helpers"
	userModel "sports-competition/app/modules/user/models"
	"time"

	"gorm.io/gorm"
)

type userDBConnection struct {
	connection *gorm.DB
}

type UserRepository interface {
	GetUserByUserID(userID string) (*userModel.User, error)
	GetUserByUsername(username string) (*userModel.User, error)
	SaveUser(createData map[string]interface{}) (*userModel.User, error)
}

func NewUserRepository() UserRepository {
	return &userDBConnection{
		connection: database.ConnectDB(),
	}
}

func (db *userDBConnection) GetUserByUserID(userID string) (*userModel.User, error) {
	var user userModel.User

	err := db.connection.Model(&user).
		Where("user_id = ?", userID).
		Find(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (db *userDBConnection) GetUserByUsername(username string) (*userModel.User, error) {
	var user userModel.User

	err := db.connection.Model(&user).
		Where("username = ?", username).
		Find(&user).
		Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *userDBConnection) SaveUser(createData map[string]interface{}) (*userModel.User, error) {
	var user userModel.User

	helpers.JsonToStruct(&createData, &user)
	var datetimeNow time.Time = helpers.GetDateTimeNow()
	user.CreatedAt = datetimeNow
	user.UpdatedAt = datetimeNow
	if user.UserID == "" {
		user.UserID = helpers.GetUniqueString()
	}

	err := db.connection.Model(&user).Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
