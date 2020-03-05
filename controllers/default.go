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
	Post_type		string`json:"post_type"`

	Request_type	string`json:"request_type"`
	Comment 		string`json:"comment"`
	Flag  			string`json:"flag"`

	Message_type	string`json:"message_type"`
	Message_id		int64`json:"message_id"`

	User_id 		int64`json:"user_id"`
	Group_id		int64`json:"group_id"`
	Discuss_id		int64`json:"discuss_id"`

	Message 		string`json:"message"`
}

type AddFriendReply struct {
	Flag			string`json:"flag"`
	Approve			bool`json:"approve"`
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
}

func reply_message(user MessageQQ) {
	urlSendMessage := "http://192.168.0.1:5700/send_msg"
	jsonReply := new(bytes.Buffer)
	json.NewEncoder(jsonReply).Encode(user)
	reps, err := http.Post(urlSendMessage, "application/json;charset=utf-8", jsonReply)
	if err != nil {
		fmt.Println("there is some err about add friend reply",  err.Error())
	} else {
		fmt.Println(reps)
	}
}

func add_friend(user MessageQQ) {
	urlAddFriend := "http://192.168.0.1:5700/set_friend_add_request"
	contentType := "application/json;charset=utf-8"
	addReply := AddFriendReply{}
	addReply.Approve = true
	addReply.Flag = user.Flag
	jsonReply := new(bytes.Buffer)
	json.NewEncoder(jsonReply).Encode(addReply)
	reps, err := http.Post(urlAddFriend, contentType, jsonReply)
	if err != nil {
		fmt.Println("there is some err about message reply",  err.Error())
	} else {
		fmt.Println(reps)
	}
}

func (this *MainController) Post() {
	user := MessageQQ{}
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	if user.Post_type == "message" {
		if user.Message_type == "private" {
			reply_message(user)
		} else if user.Message_type == "group" {
			reply_message(user)
		}
	} else if user.Post_type == "request" {
		if user.Request_type == "friend" {
			add_friend(user)
		}
	}
}
