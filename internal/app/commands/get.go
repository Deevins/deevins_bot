package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "TBDBDD")

	c.bot.Send(msg)
}
