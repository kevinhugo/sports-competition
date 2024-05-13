package helpers

import (
	"reflect"
	"strconv"
	"time"
)

func GetDateTimeNow() time.Time {
	return time.Now().UTC()
}

func GetTimestampUnixNanoNow() string {
	var unixTimestamp int64 = GetDateTimeNow().UnixNano()
	return strconv.Itoa(int(unixTimestamp))
}

/*
Adjust the given time and increment/decrement its hour based on given timezone

Example :

AdjustTimeByTimeZone("2023-10-10 12:00:00", "Asia/Jakarta")

return "2023-10-10 05:00:00"
*/
func AdjustTimeToUTC(timestamp interface{}, params ...interface{}) (time.Time, error) {

	// Get Timestamp from params first index [0], "Asia/Jakarta" if not provided
	var timezone string
	if len(params) > 0 {
		timezone = params[0].(string)
	} else {
		timezone = "Asia/Jakarta"
	}
	// \ Get Timestamp from params first index [0], "Asia/Jakarta" if not provided

	// Convert to Timestamp if input is string
	var newTime time.Time
	var err error
	switch reflect.TypeOf(timestamp).String() {
	case "string":
		newTime, err = time.Parse("2006-01-02 15:04:05", timestamp.(string))
	case "*string":
		newTime, err = time.Parse("2006-01-02 15:04:05", *(timestamp.(*string)))
	case "time.Time":
		newTime = timestamp.(time.Time)
	}
	// \ Convert to Timestamp if input is string

	// Adjust and counter the hours based on given timezone
	loc, _ := time.LoadLocation(timezone)
	_, offset := newTime.In(loc).Zone()
	newTime = newTime.Add(time.Second * -time.Duration(offset))
	// \ Adjust and counter the hours based on given timezone

	return newTime, err
}

/*
Returns UTC adjusted datetime now
*/
func GetDatetimeUTCNow() time.Time {
	datetimeUTC, _ := AdjustTimeToUTC(time.Now())
	return datetimeUTC
}

func CurrentTime() (formattedTime string) {
	loca, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().In(loca)
	layout := "2006-01-02T15:04:05+07:00"
	return currentTime.Format(layout)
}
