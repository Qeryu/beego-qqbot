package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func reply_message(user MessageQQ) {
	user.Message = "你好，" + user.Sender.Nickname + ":" + user.Message

	urlSendMessage := "http://192.168.0.1:5700/send_msg"
	contenType := "application/json;charset=utf-8"
	jsonReply := new(bytes.Buffer)
	json.NewEncoder(jsonReply).Encode(user)
	reps, err := http.Post(urlSendMessage, contenType, jsonReply)
	if err != nil {
		fmt.Println("err message reply",  err.Error())
	} else {
		fmt.Println(reps)
	}
}