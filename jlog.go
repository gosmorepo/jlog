package jlog

import (

    "fmt"
    "time"
    "encoding/json"
    "log"

)

const (

	INFO = "info"
	FATAL = "fatal"
	WARN = "warning"
	ERROR = "error"
	INIT = "Init"

)

type Fields map[string]interface{}

type logData struct {

	Timestamp string  				`json:"time"`
	Src string  			 		`json:"src"`
	Method string   				`json:"method"`
	MsgType string 					`json:"type"`
	Msg string   					`json:"message"`
	Data map[string]interface{}  	`json:"data"`

}

func Log(stdOut bool, logOut bool, src string, method string, msgType string, msg string, data map[string]interface{}) {

	var lData logData

	t := time.Now()
	lData.Timestamp = t.String()
	lData.Src = src
	lData.Method = method
	lData.MsgType = msgType
	lData.Msg = msg
	lData.Data = data

	// marshal to json
	data_json, err := json.Marshal(lData)

	data_json_s := string(data_json)

	// handle errors
	if err != nil {
		fmt.Println(err)
	}

	if stdOut {
		fmt.Println(data_json_s)
	}

	if logOut {
		log.Println(data_json_s)
	}
    
}