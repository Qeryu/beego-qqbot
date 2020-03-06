package main

import (
	"log"
	"net/http"
	"qqbotapi"
)

func main() {
	bot, err := qqbotapi.NewBotAPI("123456789ab", "http://192.168.0.1:5700", "CQHTTP_SECRET")
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

	http.ListenAndServe("0.0.0.0:8443", nil)
}
