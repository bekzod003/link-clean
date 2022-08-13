package tgbotapi

import (
	"gopkg.in/telebot.v3"

	"github.com/bekzod003/link-clean/internal/controller/tgbotapi/bot_handler"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type TelegramBot struct {
	bot      *telebot.Bot
	handlers *bot_handler.TelegramBotHandler
	log      logger.LoggerI
}

func NewTelegramBot(bot *telebot.Bot, handlers *bot_handler.TelegramBotHandler, log logger.LoggerI) *TelegramBot {
	return &TelegramBot{
		bot:      bot,
		handlers: handlers,
		log:      log,
	}
}

func (tgbot *TelegramBot) Run() {
	tgbot.log.Info("Running telegram bot")

	tgbot.bot.Handle("/start", tgbot.handlers.Start)
	tgbot.bot.Handle("/alo", tgbot.handlers.Alo)

	tgbot.log.Info("Telegram bot is starting...")
	tgbot.bot.Start()
}
