package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/singl3focus/em_testtask/config"
	"github.com/singl3focus/em_testtask/internal/http"
	"github.com/singl3focus/em_testtask/internal/http/handler"
	"github.com/singl3focus/em_testtask/internal/repo/postgres"
	"github.com/singl3focus/em_testtask/internal/service"
	mylogger "github.com/singl3focus/em_testtask/pkg/logger"
)

// @title Songs API
// @version 1.0
// @description This is a song library API as a test assignment for the company Effective mobile
// @termsOfService http://swagger.io/terms/
// @BasePath /
func main() {
	// Init configs and needed objects
	cfg := config.GetConfig("./api.env")

	loggerCfg := mylogger.NewCongig(
		cfg.Logger.LogsDirPath, cfg.Logger.Level, cfg.Logger.Format, cfg.Logger.Enable)
	logger, logFile := mylogger.SetupLogger(loggerCfg, "songs_api")
	if logger == nil {
		log.Fatal("logger not setted")
	}

	log.Printf("%#v", cfg)

	time.Sleep(15 * time.Second)

	r, err := postgres.NewPostgresDB(cfg.Database.Link, logger)
	if err != nil {
		logger.Error("connect to postgres db error", "(error)", err)
		return
	}
	s := service.NewService(r, logger)
	h := handler.NewHandler(s, logger)
	server := http.NewServer(cfg.Server.Port, h.Router())

	// App Start
	sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

    go func() {
		server.Start()
    }()

	logger.Info("Приложение запущено")

	// App Shutdown
    <-sigs
    
	logger.Info("Остановка работы приложения")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("failed to stop server", "(error)", err)
	}

    logger.Info("Работа приложения прекращена")
	mylogger.CloseLogger(logFile)
}