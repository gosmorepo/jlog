package jlog

import (

    "fmt"
    "time"
    "encoding/json"

)

const (

	INFO = "info"
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

func Log(stdOut bool, log bool, src string, method string, msgType string, msg string, data map[string]interface{}) {

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

	// handle errors
	if err != nil {
		fmt.Println(err)
	}

	if stdOut {
		fmt.Println(string(data_json))
	}
    
}