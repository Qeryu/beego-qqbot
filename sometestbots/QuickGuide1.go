package main

import(
	"net/http";
	"log";
	"qq-bot-api"
)

func main() {
	bot, err := qqbotapi.NewBotAPI("123456789ab", "http://192.168.0.1:5700", "CQHTTP_SECRET")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := qqbotapi.NewWebhook("/webhook_endpoint")
	u.PreloadUserInfo = true

	// Use WebHook as event method
	updates := bot.ListenForWebhook(u)
	// Or if you love WebSocket Reverse
	// updates := bot.ListenForWebSocket(u)
	go http.ListenAndServe("0.0.0.0:8443", nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)

		bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)
	}
}
