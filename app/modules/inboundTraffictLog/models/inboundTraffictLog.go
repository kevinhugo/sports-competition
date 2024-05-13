package models

import (
	"time"

	"gorm.io/gorm"
)

type InboundTraffictLog struct {
	gorm.Model
	ID                 uint       `json:"id" gorm:"column:id;not null:true;type:bigint;autoIncrement;primaryKey;"`
	TrackingID         string     `json:"tracking_id" gorm:"column:tracking_id;not null:true;type:VARCHAR(50);index:tracking_id_unique;unique;"`
	RequestName        string     `json:"request_name" gorm:"column:request_name;not null:false;type:VARCHAR(50)"`
	RequestHost        string     `json:"request_host" gorm:"column:request_host;not null:false;type:TEXT;"`
	RequestURL         string     `json:"request_url" gorm:"column:request_url;not null:false;type:TEXT;"`
	RequestHeader      string     `json:"request_header" gorm:"column:request_header;not null:true;TEXT;"`
	RequestBody        *string    `json:"request_body" gorm:"column:request_body;not null:false;type:TEXT;"`
	RequestTime        time.Time  `json:"request_time" gorm:"column:request_time;not null:false;type:timestamptz;"`
	ResponseStatusCode *uint      `json:"response_status_code" gorm:"column:response_status_code;not null:false;type:smallint;"`
	ResponseHeader     *string    `json:"response_header" gorm:"column:response_header;not null:false;type:TEXT;"`
	ResponseBody       *string    `json:"response_body" gorm:"column:response_body;not null:false;type:TEXT;"`
	ResponseTime       *time.Time `json:"response_time" gorm:"column:response_time;not null:false;type:timestamptz;"`
	CreatedAt          time.Time  `json:"created_at" gorm:"column:created_at;not null:true;type:timestamptz;default:current_timestamp;"`
}
