package tgbotapi

import (
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"

	"github.com/bekzod003/link-clean/internal/domain/usecase/link"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type TelegramBot struct {
	bot     *telebot.Bot
	usecase *link.UsecaseLink
	log     logger.LoggerI
}

func NewTelegramBot(bot *telebot.Bot, usecase *link.UsecaseLink, log logger.LoggerI) *TelegramBot {
	return &TelegramBot{
		bot:     bot,
		usecase: usecase,
		log:     log,
	}
}

func (tgbot *TelegramBot) Run() {
	tgbot.log.Info("Running telegram bot")

	tgbot.bot.Handle("/alo", func(c telebot.Context) error {
		tgbot.log.Info("Request to /alo from user", zap.Any("user", c.Sender()))
		if err := c.Send("Sup man, how you doing?"); err != nil {
			tgbot.log.Error("Error while sending response to /alo endpint", zap.Error(err))
			return err
		}
		return nil
	})

	tgbot.log.Info("Telegram bot is starting...")
	tgbot.bot.Start()
}
