package bot_handler

import (
	"time"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"

	"github.com/bekzod003/link-clean/internal/domain/usecase/link"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type TelegramBotHandler struct {
	usecase *link.UsecaseLink
	log     logger.LoggerI
}

func NewTelegramBotHandler(usecase *link.UsecaseLink, log logger.LoggerI) *TelegramBotHandler {
	return &TelegramBotHandler{
		usecase: usecase,
		log:     log,
	}
}

func (h *TelegramBotHandler) Alo(c telebot.Context) error {
	h.log.Info("Request to /alo from user", zap.Any("user", c.Sender()))
	if err := c.Send("Sup man, how you doing?"); err != nil {
		h.log.Error("Error while sending response to /alo endpoint", zap.Error(err))
		return err
	}
	return nil
}

func (h *TelegramBotHandler) Start(c telebot.Context) error {
	h.log.Info("Start from user", zap.Int64("id", c.Sender().ID))
	user, err := h.usecase.CreateUser(&link.User{
		ID:        c.Sender().ID,
		Username:  c.Sender().Username,
		FirstName: c.Sender().FirstName,
		LastName:  c.Sender().LastName,
		CreatedAt: time.Now(),
	})
	if err != nil {
		h.log.Error("Error while creating user", zap.Error(err))
		return err
	}

	h.log.Info("Successfully has been created user", zap.Any("user", user))
	// TODO: greeting message with buttons)

	c.Send("Hello, this is greeting message!")
	return nil
}
