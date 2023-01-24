package main

import (
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

	for update := range updates {
		if update.Message == nil { // If we did not get a message
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)

		case "list":
			listCommand(bot, update.Message, productService)
			// 2:09:50
		default:
			defaultBehaviour(bot, update.Message)
		}

	}

}
func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help \n"+
			"/list - list products",
	)
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsgText := "All the products: \n\n"

	products := productService.List()
	for _, p := range *products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	bot.Send(msg)
}

func defaultBehaviour(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote "+inputMessage.Text)

	bot.Send(msg)
}
