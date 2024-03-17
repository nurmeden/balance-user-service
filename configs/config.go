package configs

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

var (
	configFileNotFoundError viper.ConfigFileNotFoundError
	configParseError        viper.ConfigParseError
)

type Config struct {
	database Database
	server   Server
	logger   Logger
}

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	SSLMode  bool   `env:"DB_SSL_MODE"`
}

type Server struct {
	Port string `env:"SERVER_PORT"`
}

type Logger struct {
	Level string `env:"LOG_LEVEL"`
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		if errors.As(err, &configParseError) {
			return nil, errors.New("config parse error")
		}
		log.Printf("unable to decode into struct, %v", err)
		return nil, errors.New("")
	}

	return &c, err
}
