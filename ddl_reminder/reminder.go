package ddl_reminder

import (
	"demo1/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"strconv"
)

func HelpDDL(user_id int64) {
	o := orm.NewOrm()
	bot, err := qqbotapi.NewBotAPI("123456789ab",
		"http://"+ beego.AppConfig.String("serverhost") +":5700", "")
	qqnum := "qq" + strconv.FormatInt(user_id, 10)

	if useDdlChecker(user_id) == false {
		//此处应该新建一个表，存储该用户的ddl信息
		sqlSentence := "CREATE TABLE " + qqnum +
			" ( `id`  INT UNSIGNED AUTO_INCREMENT, " +
			"`thing` VARCHAR(1000), " +
			"`time` VARCHAR(100), " +
			"`ad_hours` VARCHAR(100), " +
			"`ticker_time` VARCHAR(100), " +
			"PRIMARY KEY (`id`)) " +
			"ENGINE=InnoDB DEFAULT CHARSET=UTF8"
		_, createErr := o.Raw(sqlSentence).Exec()
		if createErr != nil {
			fmt.Println("error in new user table", err.Error())
		}
		//在用户表中修改状态
		o.Raw("UPDATE users SET use_ddl_reminder='Y' WHERE user_id=?",
			strconv.FormatInt(user_id, 10)).Exec()
	}
	bot.NewMessage(user_id, "private").
		Text("本垃圾的简单使用说明(请注意下文所有逗号均为英文逗号)：").
		FaceByName("可怜").NewLine().
		Text("1.回复\"add,作文,2020-02-20 19:20:20,2,30\"即可添加一个名叫\"作文的\"任务").NewLine().
		Text("该任务将在DDL前2小时，每隔30分钟提醒一次").NewLine().
		Text("2.回复\"list\"即可列出自己的所有DDL以及编号").NewLine().
		Text("3.回复\"delete,2\"即可删除编号为2的任务").NewLine().
		Text("4.回复\"help\"再次显示本帮助").NewLine().
		Text("已经设置的任务，在结束前的24小时内每小时会提醒一次。").
		Send()
}

func AddDDL(user_id int64, task string, time string, adHours string, tickerTime string) {
	o := orm.NewOrm()
	bot, _ := qqbotapi.NewBotAPI("123456789ab",
		"http://"+ beego.AppConfig.String("serverhost") +":5700", "")
	qqnum := "qq" + strconv.FormatInt(user_id, 10)

	sqlSentence := "INSERT INTO " +
		qqnum +
		" (thing, time, ad_hours, ticker_time) VALUES ('" +
		task + "', '" + time + "', '" +
		adHours + "', '" + tickerTime + "')"

	res, addErr := o.Raw(sqlSentence).Exec()
	ddlId,_ := res.LastInsertId()

	if addErr != nil {
		fmt.Println("error in add task", addErr.Error())
		bot.NewMessage(user_id, "private").
			Text("error in add task").
			Send()
		fmt.Println("err in add task", addErr.Error())
	} else {
		bot.NewMessage(user_id, "private").
			Text("Add ddl successfully Id is " + strconv.FormatInt(ddlId, 10)).
			Send()

		go AddStartClock(user_id, task, time, adHours, tickerTime,
			strconv.FormatInt(ddlId, 10))
	}
}

func ListDDL(user_id int64) {
	o := orm.NewOrm()
	bot, _ := qqbotapi.NewBotAPI("123456789ab",
		"http://"+ beego.AppConfig.String("serverhost") +":5700", "")
	qqnum := "qq" + strconv.FormatInt(user_id, 10)

	var ddls []models.Ddls
	sqlSentence := "SELECT * FROM " + qqnum
	num, listErr := o.Raw(sqlSentence).QueryRows(&ddls)
	if listErr != nil {
		fmt.Println("error in list task", listErr.Error())
		bot.NewMessage(user_id, "private").
			Text("error in add list").
			Send()
	}

	bot.NewMessage(user_id, "private").
		Text("总共： " + strconv.FormatInt(num, 10) + " 个DDL：").
		Send()

	for _, ddl := range ddls {
		bot.NewMessage(user_id, "private").
			Text(strconv.Itoa(ddl.Id) + ",").
			Text(ddl.Thing + ",").
			Text(" DDL时间是： " + ddl.Time).
			Send()
	}
}

func DeleteDDL(user_id int64, ddl_id string) {
	o := orm.NewOrm()
	bot, _ := qqbotapi.NewBotAPI("123456789ab",
		"http://"+ beego.AppConfig.String("serverhost") +":5700", "")
	qqnum := "qq" + strconv.FormatInt(user_id, 10)

	sqlSentence := "DELETE FROM " +
		qqnum +
		" WHERE id=" +
		ddl_id
	_,delErr := o.Raw(sqlSentence).Exec()
	if delErr != nil {
		fmt.Println("error in delete task", delErr.Error())
		bot.NewMessage(user_id, "private").
			Text("error in delete list").
			Send()
	} else {
		bot.NewMessage(user_id, "private").
			Text("delete task " + ddl_id + " successfully").
			Send()
	}
}

func useDdlChecker(user_id int64) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("users")
	var user models.Users
	qs.Filter("user_id", user_id).One(&user)
	if user.UseDdlReminder == "N" {
		return false
	} else {
		return true
	}
}