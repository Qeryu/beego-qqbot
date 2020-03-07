package ddl_reminder

import (
	"demo1/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"strconv"
	"time"
)

func AddStartClock(userId int64, ddlThing string, ddlTime string, adHours string, tickerTime string, ddlId string) {
	checkExist(userId, ddlId)

	withNanos := "2006-01-02 15:04:05"
	stTime,_ := time.Parse(withNanos, ddlTime)

	tickerTimeNum,_ := strconv.Atoi(tickerTime)
	adHoursNum,_ := strconv.Atoi(adHours)

	totalSecond := stTime.Sub(time.Now())- time.Duration(adHoursNum * 3600)* time.Second - 8 * time.Hour
	timer := time.NewTimer(totalSecond)
	<-timer.C

	bot, _ := qqbotapi.NewBotAPI("123456789ab",
		"http://"+ beego.AppConfig.String("serverhost") +":5700", "")

	ticker := time.NewTicker(time.Duration(tickerTimeNum) * time.Minute)
	go func() {
		for range ticker.C {
			if checkExist(userId, ddlId) == true {
				bot.NewMessage(userId, "private").
					Text("halo，你的 " + ddlId + " 号ddl：").NewLine().
					Text(ddlThing).NewLine().
					Text("要到期了，请留意").Send()
			} else {
				return
			}
			
		}
	}()
	time.Sleep(time.Hour * time.Duration(adHoursNum))
	ticker.Stop()
	bot.NewMessage(userId, "private").
		Text("halo，你的 " + ddlId + " 号ddl：").NewLine().
		Text(ddlThing).NewLine().
		Text("要到期了，请留意，本次为最后一次提醒，ddl将被删除").Send()
	DeleteDDL(userId, ddlId)
}

func checkExist(userId int64, ddlId string) bool {
	o := orm.NewOrm()
	qqnum := "qq" + strconv.FormatInt(userId, 10)
	sqlSentence := "SELECT * FROM " + qqnum + " WHERE id=" + ddlId
	var ddl []models.Ddls
	num, _ := o.Raw(sqlSentence).QueryRows(&ddl)

	if num == 0 {
		return false
	} else {
		return true
	}
}