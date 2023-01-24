package commands

import (
	"deevins_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot: bot,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	if update.Message == nil { // If we did not get a message
		return
	}

	switch update.Message.Command() {

	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)

	default:
		c.Default(update.Message)
	}
}
