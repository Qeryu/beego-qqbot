package controllers

import (
	"demo1/ddl_reminder"
	"strings"
)

type MessageQQ struct {
	Post_type		string`json:"post_type"`

	Request_type	string`json:"request_type"`
	Comment 		string`json:"comment"`
	Flag  			string`json:"flag"`

	Message_type	string`json:"message_type"`
	Message_id		int64`json:"message_id"`

	User_id 		int64`json:"user_id"`
	Sender			*Sender`json:"sender"`

	Group_id		int64`json:"group_id"`
	Discuss_id		int64`json:"discuss_id"`

	Message 		string`json:"message"`
	Raw_message		string`json:"raw_message"`
}

type Sender struct {
	Nickname	string`json:"nickname"`
	Sex			string`json:"sex"`
	Age			uint8`json:"age"`
}

func deal_messages(user MessageQQ) {
	if user.Post_type == "message" {
		if user.Message_type == "private" {
			if ddlChecker(user) == false {
				reply_message(user)
			}
		} else if user.Message_type == "group" {
			reply_message(user)
		}
	} else if user.Post_type == "request" {
		if user.Request_type == "friend" {
			add_friend(user)
		}
	}
}

func ddlChecker(user MessageQQ) bool {
	ddlCommand := strings.Split(user.Raw_message, ",")

	if ddlCommand[0] == "help" {
		ddl_reminder.HelpDDL(user.User_id)
	} else if ddlCommand[0] == "add" {
		ddl_reminder.AddDDL(user.User_id, ddlCommand[1], ddlCommand[2], ddlCommand[3], ddlCommand[4])
	} else if ddlCommand[0] == "list" {
		ddl_reminder.ListDDL(user.User_id)
	} else if ddlCommand[0] == "delete" {
		ddl_reminder.DeleteDDL(user.User_id, ddlCommand[1])
	} else {
		return false
	}
	return true
}