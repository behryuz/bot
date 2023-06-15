package main

import (
	"github.com/behryuz/bot/internal/app/commands"
	"github.com/behryuz/bot/internal/service/product"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "list":
			commander.List(update.Message, productService)
		default:
			commander.Default(update.Message)
		}
	}
}
