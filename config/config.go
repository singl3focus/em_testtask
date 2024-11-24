package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Logger   Logger   `yaml:"logger"`
}

type Server struct {
	Port string `yaml:"port" env-default:"5005"`
}

type Database struct {
	Link string `yaml:"link" env-required:"true" env:"DATABASE_CONNECT_LINK"`
}

type Logger struct {
	LogsDirPath string `yaml:"logs_dir" env-default:"./logs" env:"LOGGER_DIR"`
	Enable      bool   `yaml:"enable" env-default:"true" env:"LOGGER_ENABLE"`
	Level       string `yaml:"level" env-default:"DEBUG" env:"LOGGER_LEVEL"`
	Format      string `yaml:"format" env-default:"TXT" env:"LOGGER_FORMAT"`
}

func GetConfig(configPath string) *Config {
	_, err := os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("config file does not exist: %s", configPath)
		}
		log.Fatal(err)
	}

	var cfg *Config
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}
