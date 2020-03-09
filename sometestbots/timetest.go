package main

import (
	"fmt"
	"github.com/astaxie/beego"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"strconv"
	"time"
)

func AddStartClock(userId int64, ddlThing string, ddlTime string, adHours string, tickerTime string, ddlId string) {
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
			if true {
				bot.NewMessage(userId, "private").
					Text("halo，你的 " + ddlId + " 号ddl：").NewLine().
					Text(ddlThing).NewLine().
					Text("要到期了，请留意").Send()
			}
		}
	}()

	if time.Hour * time.Duration(adHoursNum) > stTime.Sub(time.Now()) - 8 * time.Hour {
		time.Sleep(stTime.Sub(time.Now()) - 8 * time.Hour)
	} else {
		time.Sleep(time.Hour * time.Duration(adHoursNum))
	}
	ticker.Stop()

	if true {
		bot.NewMessage(userId, "private").
			Text("halo，你的 " + ddlId + " 号ddl：").NewLine().
			Text(ddlThing).NewLine().
			Text("要到期了，请留意，本次为最后一次提醒，ddl将被删除").Send()
		fmt.Println("delete over")
	}
}

func main() {
	AddStartClock(1421683965, "task", "2020-03-09 08:08:00", "1", "1", "1")
}