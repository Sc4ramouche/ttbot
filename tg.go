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

	// it's super naive to just use hardcoded id, but for now it serves the
	// purpose. Ideally, I should store the list of users who had interacted
	// with the bot in database.
	// SQLite could be a good fit here.
	msg := tgbotapi.NewMessage(id, "The enrollment link is available:\nhttps://sttkberlin.de/trainingszeiten/")
	bot.Send(msg)
}
