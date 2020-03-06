package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type AddFriendReply struct {
	Flag			string`json:"flag"`
	Approve			bool`json:"approve"`
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
}

func (this *MainController) Post() {
	user := MessageQQ{}
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	deal_messages(user)
}
