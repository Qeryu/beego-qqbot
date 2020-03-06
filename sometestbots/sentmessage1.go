package main

import (
	"github.com/catsworld/qq-bot-api"
	"log"
	//"net/http"
)

func main() {
	bot, err := qqbotapi.NewBotAPI("123456789ab", "http://175.24.23.211:5700", "")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := qqbotapi.NewWebhook("/webhook_endpoint")
	u.PreloadUserInfo = true

	bot.ListenForWebhookSync(u, func(update qqbotapi.Update) interface{} {

		log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)

		return map[string]interface{}{
			"reply": update.Message.Text,
		}
	})

	// Send a text-img message
	s :=bot.NewMessage(1421683965, "private").
		At("1232332333").
		Text("嘤嘤嘤").
		NewLine().
		FaceByName("可怜").
		Text("这是一个测试").
		ImageBase64("img.jpg").
		Send()

	//Withdraw that message
	if s.Err == nil {
		bot.DeleteMessage(s.Result.MessageID)
	}

	// Send a stand-alone message (No need to call Send())
	bot.NewMessage(10000000, "private").
		Dice()
}
