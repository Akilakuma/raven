package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	ChatId   = -111111
	BotToken = "your token"
)

var Bot *tgbotapi.BotAPI

func newBot() {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	Bot = bot

	// telegram bot receive msg
	go recvMsg()
}

//  recvMsg log polling 方式持續看telegram bot 收到的訊息
func recvMsg() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		Bot.Send(msg)
	}
}

func sendFile(filename string) {
	msg := tgbotapi.NewDocumentUpload(ChatId, filename)
	_, err := Bot.Send(msg)

	if err != nil {
		log.Println("🐯 sendFile error", err)
	}
}

func sendPhoto(filename string) {
	msg := tgbotapi.NewPhotoUpload(ChatId, filename)
	_, err := Bot.Send(msg)

	if err != nil {
		log.Println("🐯 sendFile error", err)
	}
}
