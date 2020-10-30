package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const ChatId = -12345677

var Bot *tgbotapi.BotAPI

func main() {

	// 啟動telegram bot
	newBot()

	// 啟動gin
	r := gin.New()
	r.POST("api/file-upload", fileUpload)
	r.Run(":8866")
}

func newBot() {
	bot, err := tgbotapi.NewBotAPI("your telegram token")
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

func fileUpload(c *gin.Context) {
	// request 參數
	file, _, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// 新建立一個空的file，放在output
	output, err := os.Create("report.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	// 將從request拿到的file，copy到output
	_, err = io.Copy(output, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful \n")

	go sendFile("report.csv")
}
