package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

// Status Enum
type UserStatus string

const (
	USER_STATUS_ACTIVE   UserStatus = "ACTIVE"
	USER_STATUS_INACTIVE UserStatus = "INACTIVE"
)

var userStatusList []UserStatus = []UserStatus{
	USER_STATUS_ACTIVE,
	USER_STATUS_INACTIVE,
}

func GetUserStatusList() []UserStatus {
	var statusList []UserStatus = make([]UserStatus, len(userStatusList))
	copy(statusList, userStatusList)
	return statusList
}

func (e *UserStatus) Scan(value interface{}) error {
	switch value.(type) {
	case []byte:
		break
	default:
		value = []byte(value.(string))
	}
	*e = UserStatus(value.([]byte))
	return nil
}

func (e UserStatus) UserStatus() (driver.Value, error) {
	return string(e), nil
}

func MigrateUserStatusEnum(db *gorm.DB) {
	db.Exec(`
		CREATE TYPE user_status AS ENUM (
			'ACTIVE',
			'INACTIVE'
		);
	`)
}

// \ Status Enum

type User struct {
	gorm.Model
	ID        uint       `json:"id" gorm:"column:id;not null:true;type:bigint;autoIncrement;primaryKey;"`
	UserID    string     `json:"user_id" gorm:"column:user_id;not null:true;type:VARCHAR(50);index:user_id_unique;unique;"`
	Username  string     `json:"username" gorm:"column:username;not null:true;type:VARCHAR(50);"`
	Password  string     `json:"password" gorm:"column:password;not null:true;type:TEXT;"`
	Status    UserStatus `json:"status" sql:"type:user_status;" gorm:"column:status;type:user_status;not null:true;default:'ACTIVE'"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;not null:true;type:timestamptz;default:current_timestamp;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;not null:true;type:timestamptz;default:current_timestamp;"`
}
