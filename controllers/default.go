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
	User_id 		int64 `json:"user_id"`
	Message 		string `json:"message"`
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func reply_message(user MessageQQ) {
	url := "http://192.168.0.1:5700/send_private_msg"
	jsonReply := new(bytes.Buffer)
	json.NewEncoder(jsonReply).Encode(user)
	http.Post(url, "application/json;charset=utf-8", jsonReply)
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
