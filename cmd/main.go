package main

import (
	"github.com/nurmeden/balance-user-service/configs"
	"github.com/nurmeden/balance-user-service/pkg/utils"
	"log"
	"os"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := configs.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	config, err := configs.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Ошибка при парсинге конфигурации: %v", err)
	}

	startServer(config)
}

func startServer(config *configs.Config) {
}
