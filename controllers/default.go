package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"net/http"
)

type MainController struct {
	beego.Controller
}


type MessageQQ struct {
	post_type 		string
	usr_id 			string
	message 		string
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func reply_message(user MessageQQ) {
	url := "http://192.168.0.1:5700/send_private_msg"
	jsonReply,_ := json.Marshal(user)
	http.NewRequest("POST", url, bytes.NewBuffer(jsonReply))
	fmt.Println("ok")
}

func (this *MainController) Post() {
	var user MessageQQ
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(user)
	if user.post_type == "message" {
		reply_message(user)
	}
}
