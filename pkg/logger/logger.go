package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	LogsDirPath string
	Enable   bool
	Level    string
	Format   string
}

func NewCongig(logsDirPath, level, format string, enable bool) *Config {
	return &Config{
		LogsDirPath: logsDirPath,
		Level: level,
		Format: format,
		Enable: enable,
	}
}

/*	
	SetupLogger.
		Avaliable levels: DEBUG, INFO, WARN, ERROR. Default: DEBUG
		Avaliable formats: JSON, TXT. Default: TXT
	Return:
		logger(*slog.Logger) - if it is created
		path to log file (*os.File)
*/
func SetupLogger(cfg *Config, apiName string) (*slog.Logger, *os.File) {
	if !cfg.Enable {
		return nil, nil
	}

	var level slog.Level

	switch strings.ToUpper(cfg.Level) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN", "WARNING":
		level = slog.LevelWarn
	case "ERROR", "ERR":
		level = slog.LevelError
	default:
		level = slog.LevelDebug	
	}


	var logger *slog.Logger
	var logsFile *os.File

	switch strings.ToUpper(cfg.Format) {
	case "JSON":
		filePath := strings.ReplaceAll(fmt.Sprintf("%s/%s.json", cfg.LogsDirPath, apiName), "//", "/")
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		
		logger = slog.New(
			slog.NewJSONHandler(file, &slog.HandlerOptions{Level: level}),
		)

		logsFile = file
	default:
		filePath := strings.ReplaceAll(fmt.Sprintf("%s/%s.log", cfg.LogsDirPath, apiName), "//", "/")
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		logger = slog.New(
			slog.NewTextHandler(file, &slog.HandlerOptions{Level: level}),
		)

		logsFile = file
	}

	return logger, logsFile
}

func CloseLogger(logsFile *os.File) {
	err := logsFile.Close()
	if err != nil {
		log.Fatalf("close log file error: %e", err)
	}
}