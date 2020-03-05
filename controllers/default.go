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
	Group_id		string`json:"group_id"`
	Discuss_id		string`json:"discuss_id"`

	Message 		string`json:"message"`
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func reply_message(user MessageQQ) {
	urlSendMessage := "http://175.24.23.211:5700/send_group_msg"
	jsonReply := new(bytes.Buffer)
	json.NewEncoder(jsonReply).Encode(user)
	reps, err := http.Post(urlSendMessage, "application/json;charset=utf-8", jsonReply)
	if err != nil {
		fmt.Println("there is some err about reply")
	}
	fmt.Println(reps.Body)
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
