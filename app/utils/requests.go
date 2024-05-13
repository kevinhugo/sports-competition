package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sports-competition/app/env"
	"sports-competition/app/helpers"
	"sports-competition/app/logger"
	outboundTraffictLog "sports-competition/app/modules/outboundTraffictLog/repositories"
	"strconv"
	"time"
)

var outboundTraffictLogRepository outboundTraffictLog.OutboundTraffictLogRepository = outboundTraffictLog.NewOutboundTraffictLogRepository()
var saveOutboundLogProcessTimeout int = env.OUTBOUND_CONNECTION_TIMEOUT_IN_SECOND * 2

type RequestResponse struct {
	StatusCode uint
	Header     map[string]interface{}
	Body       map[string]interface{}
	ByteBody   []byte
}

type Request struct {
	RequestName string
}

type RequestInterface interface {
}

/*
General Request Wrapper

method param is presumed the request method

url param is presumed url destination of the request

additionalData[0] param is presumed request header

additionalData[1] param is presumed request body
*/
func (request Request) HttpRequest(method string, url string, additionalData ...interface{}) (*RequestResponse, error) {
	logger.Debug(fmt.Sprintf("------- %s request %s -------", method, request.RequestName))
	client := &http.Client{
		Timeout: time.Duration(env.OUTBOUND_CONNECTION_TIMEOUT_IN_SECOND) * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header["Content-Type"] = []string{"application/json"}
	var additionalDataLength int = len(additionalData)
	if additionalDataLength > 0 {
		if additionalData[0] != nil {
			for index, each := range additionalData[0].(map[string]interface{}) {
				req.Header[index] = []string{each.(string)}
			}
		}
	}
	var headerStringData string = helpers.ToJson(helpers.StructToMap(req.Header))

	var requestBody any
	if additionalDataLength > 1 {
		requestBody = additionalData[1]
		marshalResult, _ := json.Marshal(&requestBody)
		req.Body = io.NopCloser(bytes.NewReader(marshalResult))
	}
	var bodyStringData string = helpers.ToJson(requestBody)

	logger.Debug(fmt.Sprintf("%s Request URL : %s", method, url))
	logger.Debug(fmt.Sprintf("%s Request Header : %s", method, headerStringData))
	logger.Debug(fmt.Sprintf("%s Request Body : %s", method, bodyStringData))

	var requestTime = helpers.GetDateTimeNow()

	var bodyStringSave *string
	if bodyStringData != "" {
		bodyStringSave = &bodyStringData
	}

	var updateLogData chan map[string]interface{} = make(chan map[string]interface{})

	go request.saveOutboundTraffictLog(map[string]interface{}{
		"request_name":   request.RequestName,
		"request_url":    url,
		"request_header": headerStringData,
		"request_body":   bodyStringSave,
		"request_time":   requestTime,
	}, updateLogData)

	resp, err := client.Do(req)
	if err != nil {
		updateLogData <- map[string]interface{}{}
		return nil, err
	}
	defer resp.Body.Close()
	var responseTime = helpers.GetDateTimeNow()

	var requestResponse RequestResponse = RequestResponse{
		StatusCode: uint(resp.StatusCode),
		Header:     helpers.StructToMap(&resp.Header),
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	requestResponse.ByteBody = respBody
	var respStringBody string
	err = json.Unmarshal(respBody, &requestResponse.Body)
	if err == nil {
		respStringBody = helpers.ToJson(requestResponse.Body)
	} else {
		respStringBody = string(respBody)
	}

	var responseHeaderStringData string = helpers.ToJson(helpers.StructToMap(resp.Header))

	updateLogData <- map[string]interface{}{
		"response_status_code": resp.StatusCode,
		"response_header":      responseHeaderStringData,
		"response_body":        respStringBody,
		"response_time":        responseTime,
	}

	logger.Debug(fmt.Sprintf("%s Response Code : %s", method, strconv.Itoa(resp.StatusCode)))
	logger.Debug(fmt.Sprintf("%s Response Header : %s", method, resp.Header))
	logger.Debug(fmt.Sprintf("%s Response Body %s", method, respBody))
	logger.Debug(fmt.Sprintf("------- End of %s request %s -------", method, request.RequestName))

	return &requestResponse, nil

}

func (request *Request) saveOutboundTraffictLog(logData map[string]interface{}, updateLogData <-chan map[string]interface{}) {
	newLogData, err := outboundTraffictLogRepository.SaveOutboundTraffictLog(logData)
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
		outboundTraffictLogRepository.UpdateOutboundTraffictLog(newLogData, updatedLogData)
	case <-time.After(time.Duration(saveOutboundLogProcessTimeout) * time.Second):
		fmt.Println(fmt.Sprintf("Max timeout %s second(s) reached for trying to update log data of %s.", strconv.Itoa(saveOutboundLogProcessTimeout), newLogData))
	}
}

/*
Shorthand for HttpRequest("POST", .....)

url param is presumed url destination of the request

additionalData[0] param is presumed request header

additionalData[1] param is presumed request body
*/
func (request *Request) Post(url string, additionalData ...interface{}) (*RequestResponse, error) {
	response, err := request.HttpRequest("POST", url, additionalData...)
	return response, err
}

func (request *Request) POST(url string, additionalData ...interface{}) (*RequestResponse, error) {
	return request.Post(url, additionalData...)
}

/*
Shorthand for HttpRequest("GET", .....)

url param is presumed url destination of the request

additionalData[0] param is presumed request header

additionalData[1] param is presumed request body
*/
func (request *Request) Get(url string, additionalData ...interface{}) (*RequestResponse, error) {
	response, err := request.HttpRequest("GET", url, additionalData...)
	return response, err
}

func (request *Request) GET(url string, additionalData ...interface{}) (*RequestResponse, error) {
	return request.Get(url, additionalData...)
}

/*
Shorthand for HttpRequest("PUT", .....)

url param is presumed url destination of the request

additionalData[0] param is presumed request header

additionalData[1] param is presumed request body
*/
func (request *Request) Put(url string, additionalData ...interface{}) (*RequestResponse, error) {
	response, err := request.HttpRequest("PUT", url, additionalData...)
	return response, err
}

func (request *Request) PUT(url string, additionalData ...interface{}) (*RequestResponse, error) {
	return request.Put(url, additionalData...)
}
