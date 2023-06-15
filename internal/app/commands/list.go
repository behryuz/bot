package commands

import (
	"github.com/behryuz/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsq := "Here all of the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMsq += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsq)
	c.bot.Send(msg)
}