package models

import (
	"time"

	userModel "sports-competition/app/modules/user/models"

	"gorm.io/gorm"
)

type Sport struct {
	gorm.Model
	ID          uint `json:"id" gorm:"column:id;not null:true;type:bigint;autoIncrement;primaryKey;"`
	UserID      uint `json:"user_id" gorm:"column:user_id;not null:true;type:BIGINT;index:sport_user_id_unique;unique;"`
	User        userModel.User
	Proficiency uint      `json:"proficiency" gorm:"column:proficiency;not null:true;type:INTEGER;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null:true;type:timestamptz;default:current_timestamp;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null:true;type:timestamptz;default:current_timestamp;"`
}
