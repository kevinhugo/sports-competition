package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sports-competition/app/env"
	"sports-competition/app/helpers"
	"sports-competition/app/logger"
	"strconv"
	"time"

	inboundTraffictLog "sports-competition/app/modules/inboundTraffictLog/services"

	"github.com/gin-gonic/gin"
)

var inboundTraffictLogService *inboundTraffictLog.InboundTraffictLogService = inboundTraffictLog.NewInboundTraffictLogService()

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

/*
General logger for Inbound API Hit

additionalData[0] ( bool : default true) is presumed if function will print log

additionalData[1] ( bool : default true ) is presumed if function will record the inbound request to database
*/
func RequestLogger(APIName string, additionalData ...any) gin.HandlerFunc {
	return func(c *gin.Context) {
		var additionalDataLength int = len(additionalData)
		var printLog bool = true
		var saveInboundTraffictLog bool = true
		if additionalDataLength > 0 {
			printLog = additionalData[0].(bool)
		}
		if additionalDataLength > 1 {
			saveInboundTraffictLog = additionalData[1].(bool)
		}
		headerMapData := helpers.ToJson(helpers.StructToMap(c.Request.Header))
		if printLog {
			logger.Debug(fmt.Sprintf("========================================== %s ==========================================", APIName))
			logger.Debug(fmt.Sprintf("Request Host : %s", c.Request.Host))
			logger.Debug(fmt.Sprintf("Request URL : %s", c.Request.RequestURI))
			logger.Debug(fmt.Sprintf("Request Header : %s", headerMapData))
		}
		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error(err)
		}
		var requestBodyString string = ""
		if len(requestBody) != 0 {
			requestBodyString = helpers.MinifyJSON(requestBody)
		}
		logger.Debug(fmt.Sprintf("Request Body : %s", requestBodyString))
		var requestTime = helpers.GetDateTimeNow()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		var updateLogData chan map[string]interface{} = make(chan map[string]interface{})
		if saveInboundTraffictLog {
			go requestLoggerSaveInboundTraffictLog(map[string]interface{}{
				"request_name":   APIName,
				"request_host":   c.Request.Host,
				"request_url":    c.Request.RequestURI,
				"request_header": headerMapData,
				"request_body":   requestBody,
				"request_time":   requestTime,
			}, updateLogData)

		}

		responseWriter := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = responseWriter
		c.Next()

		var responseTime = helpers.GetDateTimeNow()
		var responseHeaderMap = helpers.ToJson(helpers.StructToMap(responseWriter.Header()))

		if saveInboundTraffictLog {
			updateLogData <- map[string]interface{}{
				"response_status_code": c.Writer.Status(),
				"response_header":      responseHeaderMap,
				"response_body":        responseWriter.body.String(),
				"response_time":        responseTime,
			}
		}

		if printLog {
			logger.Debug(fmt.Sprintf("Response Status : %s", strconv.Itoa(c.Writer.Status())))
			logger.Debug(fmt.Sprintf("Response Header : %s", responseHeaderMap))
			logger.Debug(fmt.Sprintf("Response Body : %s", responseWriter.body.String()))
			logger.Debug(fmt.Sprintf("========================================== End of %s ==========================================", APIName))
		}
	}
}

func requestLoggerSaveInboundTraffictLog(logData map[string]interface{}, updateLogData <-chan map[string]interface{}) {
	newLogData, err := inboundTraffictLogService.SaveInboundTraffictLog(logData)
	if err != nil {
		logger.Error("Error while trying to save outbound traffict log data, see logs below.")
		logger.Error(err)
		return
	}

	select {
	case updatedLogData := <-updateLogData:
		if len(updatedLogData) == 0 {
			return
		}
		inboundTraffictLogService.UpdateInboundTraffictLog(newLogData, updatedLogData)
	case <-time.After(time.Duration(env.INBOUND_CONNECTION_TIMEOUT_IN_SECOND) * time.Second):
		fmt.Println(fmt.Sprintf("Max timeout %s second(s) reached for trying to update log data of %s.", env.INBOUND_CONNECTION_TIMEOUT_IN_SECOND, newLogData))
	}
}
