package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type MainController struct {
	beego.Controller
}


type MessageQQ struct {
	Message_type	string`json:"message_type"`

	User_id 		int64`json:"user_id"`
	Group_id		int64`json:"group_id"`
	Discuss_id		int64`json:"discuss_id"`

	Message 		string`json:"message"`
}


func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
}

func reply_message(user MessageQQ) {
	urlSendMessage := "http://198.162.0.1:5700/send_msg"
	jsonReply := new(bytes.Buffer)
	json.NewEncoder(jsonReply).Encode(user)
	reps, err := http.Post(urlSendMessage, "application/json;charset=utf-8", jsonReply)
	if err != nil {
		fmt.Println("there is some err about reply")
	}
	fmt.Println(reps)
}

func (this *MainController) Post() {
	user := MessageQQ{}
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	reply_message(user)
}
