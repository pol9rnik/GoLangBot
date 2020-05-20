package main

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"os"
)
const(
	webHook = ""
)
func main() {
	port:=os.Getenv("Port")
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	bot, err := tgbotapi.NewBotAPI("1220156233:AAHIfLQS7nzdwWdt6-y1hwaIg5Zd4seMKq4")
	if err != nil{
		log.Fatal("creation bot: ", err)
	}
	log.Println("bot created")

	if _, err :=bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatalf("setting webHook %v; error: %v", webHook, err)
	}
	log.Println("webHook set")

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil{
			log.Print(err)
		}
	}

}
