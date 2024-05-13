package logger

import (
	"fmt"
	"sports-competition/app/env"
	"time"
)

/*
Basic log printer with tag and timestamp
*/
func PrintLog(tag string, message any) {
	fmt.Println(fmt.Sprintf("[%s][%s] %s", tag, time.Now().Format("2006-01-02 15:04:05.999999999 -0700"), message))
}

/*
PrintLog with DEBUG tag
*/
func Debug(message any) {
	if env.LOGGER_SHOW_DEBUG {
		PrintLog("DEBUG", message)
	}
}

/*
PrintLog with INFO tag
*/
func Info(message any) {
	if env.LOGGER_SHOW_INFO {
		PrintLog("INFO", message)
	}
}

/*
PrintLog with WARNING tag
*/
func Warning(message any) {
	if env.LOGGER_SHOW_WARNING {
		PrintLog("WARNING", message)
	}
}

/*
PrintLog with ERROR tag
*/
func Error(message any) {
	if env.LOGGER_SHOW_ERROR {
		PrintLog("ERROR", message)
	}
}
