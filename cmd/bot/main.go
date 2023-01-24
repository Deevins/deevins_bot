package main

import (
	"deevins_bot/internal/app/commands"
	"deevins_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	botToken := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil { // If we did not get a message
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)

		case "list":
			commander.List(update.Message)
			// 2:09:50
		default:
			commander.Default(update.Message)
		}

	}

}
