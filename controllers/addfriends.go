package controllers

import (
	"bytes"
	"demo1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
)

func add_friend(user MessageQQ) {
	urlAddFriend := "http://"+ beego.AppConfig.String("serverhost") +":5700/set_friend_add_request"
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

	o := orm.NewOrm()
	o.Using("default")
	addUserNewFriend(user, o)

}

func addUserNewFriend(val MessageQQ, o orm.Ormer) {
	user := new(models.Users)
	user.Nickname = val.Sender.Nickname
	user.UserId = val.User_id
	created, id, er := o.ReadOrCreate(user, "UserId")
	if er == nil {
		if created {
			fmt.Println("New insert id", id)
		} else {
			fmt.Println("Get an object", id)
		}
	} else {
		fmt.Println("some error", er.Error())
	}
}