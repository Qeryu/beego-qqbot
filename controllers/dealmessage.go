package controllers

import (
	"demo1/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"strconv"
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
	o := orm.NewOrm()
	bot, err := qqbotapi.NewBotAPI("123456789ab", "http://175.24.23.211:5700", "")

	var sqlSentence string
	qqnum := strconv.FormatInt(user.User_id, 10)
	if user.Message[0:4] == "help" {
		if useDdlChecker(user.User_id) == false {
			//此处应该新建一个表，存储该用户的ddl信息
			sqlSentence = "CREATE TABLE" + qqnum +
				"( `id`  INT UNSIGNED AUTO_INCREMENT, " +
				"`thing VARCHAR(1000), " +
				"`time` VARCHAR(100)" +
				"PRIMARY KEY (`id`)) " +
				"ENGINE=InnoDB DEFAULT CHARSET=UTF8"
			_, createErr := o.Raw(sqlSentence).Exec()
			if createErr != nil {
				fmt.Println("error in new user table", err.Error())
			}
		}
		bot.NewMessage(user.User_id, "private").
			Text("本垃圾的简单使用说明(请注意下文所有逗号均为英文逗号)：").
			FaceByName("可怜").NewLine().
			Text("1.回复\"add,作文,2020-02-20 19:20:20\"即可添加一个名叫\"作文的\"任务").NewLine().
			Text("2.回复\"list\"即可列出自己的所有DDL以及编号").NewLine().
			Text("3.回复\"delete,2\"即可删除编号为2的任务").NewLine().
			Text("4.回复\"help\"再次显示本帮助").NewLine().
			Text("已经设置的任务，在结束前的24小时内每小时会提醒一次。").
			Send()
	} else if user.Message[0:3] == "add" {
		addCommand := strings.Split(user.Message, ",")
		_, addErr := o.Raw("INSERT INTO ? (thing, time) VALUES (?, ?)",
			qqnum, addCommand[1], addCommand[2]).Exec()

		if addErr != nil {
			fmt.Println("error in add task", addErr.Error())
			bot.NewMessage(user.User_id, "private").
				Text("error in add task").
				Send()
		} else {
			bot.NewMessage(user.User_id, "private").
				Text("Add ddl successfully").
				Send()
		}

	} else if user.Message[0:4] == "list" {
		var ddls []*models.Ddls
		num, listErr := o.QueryTable(qqnum).All(&ddls)
		if listErr != nil {
			fmt.Println("error in list task", listErr.Error())
			bot.NewMessage(user.User_id, "private").
				Text("error in add list").
				Send()
		}

		bot.NewMessage(user.User_id, "private").
			Text("总共： " + strconv.FormatInt(num, 10) + " 个DDL：")

		for _, ddl := range ddls {
			bot.NewMessage(user.User_id, "private").
				Text(strconv.Itoa(ddl.Id) + ",").
				Text(ddl.Thing + ",").
				Text(ddl.Time + ",").
				Send()
		}
	} else if user.Message[0:6] == "delete" {
		deleteCommand := strings.Split(user.Message, ",")
		_,delErr := o.Raw("DELETE FROM ? WHERE id=?", qqnum, deleteCommand[1]).Exec()
		if delErr != nil {
			fmt.Println("error in delete task", delErr.Error())
			bot.NewMessage(user.User_id, "private").
				Text("error in delete list").
				Send()
		} else {
			bot.NewMessage(user.User_id, "private").
				Text("delete task " + deleteCommand[1] + " successfully").
				Send()
		}
	} else {
		return false
	}
	return true
}

func useDdlChecker(user_id int64) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("users")
	var user models.Users
	qs.Filter("user_id", user_id).One(&user)

	if user.UseDdlReminder == "N" {
		return false
	}
	return true
}