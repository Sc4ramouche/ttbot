package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

func ConnectToTGBot() {
	id, err := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	msg := tgbotapi.NewMessage(id, "The enrollment link is available:\nhttps://sttkberlin.de/trainingszeiten/")
	bot.Send(msg)
}
