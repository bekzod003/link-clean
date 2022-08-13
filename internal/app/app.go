package app

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"

	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/internal/adapters/db/postgres"
	"github.com/bekzod003/link-clean/internal/controller/tgbotapi"
	"github.com/bekzod003/link-clean/internal/controller/tgbotapi/bot_handler"
	"github.com/bekzod003/link-clean/internal/domain/service"
	"github.com/bekzod003/link-clean/internal/domain/usecase/link"
	"github.com/bekzod003/link-clean/pkg/database/client/postgresql"
	"github.com/bekzod003/link-clean/pkg/logger"
)

func Run(cfg *config.Config) {
	println("Logger initializing...")
	log := logger.NewLogger(cfg.ServiceName, cfg.LoggerLevel)
	println("Logger initialized successfully")
	defer logger.Cleanup(log)

	psqlConfig := getPostgresClientConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	poolConnection, err := postgresql.NewClient(ctx, psqlConfig)
	if err != nil {
		log.Fatal("Error while getting postgres new client", zap.Error(err))
	}
	defer func() {
		log.Info("Closing pool connection")
		poolConnection.Close()
	}()

	linkStorage := postgres.NewLinkStorage(poolConnection)
	userStorage := postgres.NewUserStorage(poolConnection)
	tagStorage := postgres.NewTagStorage(poolConnection)

	linkService := service.NewLinkService(linkStorage, log)
	tagService := service.NewTagService(tagStorage, log)
	userService := service.NewUserService(userStorage, log)

	useCase := link.NewLinkUsecase(link.ConstructorRequest{
		LinkService: linkService,
		TagService:  tagService,
		UserService: userService,
		Log:         log,
	})

	botSettings := telebot.Settings{
		Token:  cfg.Telegram.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, c telebot.Context) {
			log.Error("Error in telegram bot", zap.Error(err))
			c.Send("Error has been occurred: " + err.Error())
		},
	}

	bot, err := telebot.NewBot(botSettings)
	if err != nil {
		log.Fatal("Error while getting new bot", zap.Error(err))
	}
	botHandlers := bot_handler.NewTelegramBotHandler(useCase, log)
	tgbotapi.NewTelegramBot(bot, botHandlers, log).Run()
}

func getPostgresClientConfig(cfg *config.Config) postgresql.ClientConfig {
	return postgresql.ClientConfig{
		Login:    cfg.PostgreSQL.User,
		Password: cfg.PostgreSQL.Password,
		Host:     cfg.PostgreSQL.Host,
		Port:     cfg.PostgreSQL.Port,
		DBName:   cfg.PostgreSQL.DBName,
		PoolConfig: postgresql.PoolConfig{
			MaxConns:                 cfg.PostgreSQL.PoolConfig.MaxConns,
			MaxConnIdleMinutes:       cfg.PostgreSQL.PoolConfig.MaxConnIdleMinute,
			MaxConnLifetimeMinutes:   cfg.PostgreSQL.PoolConfig.MaxConnLifetimeMinute,
			HealthCheckPeriodMinutes: cfg.PostgreSQL.PoolConfig.HealthCheckPeriodMinute,
		},
	}
}
