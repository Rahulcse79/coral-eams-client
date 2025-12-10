package main

import (
	"encoding/json"
	"fmt"
	"os"
	"coral-eams-client/internal/logger"
)

type Config struct {
	LogFileName string `json:"logFileName"`
	Port        int    `json:"port"`
	HostName    string `json:"hostName"`
}

func loadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	logger.InitLogger(config.LogFileName)

	if config.LogFileName != "" {
		fmt.Printf("Coral EAMS Client Started and log file created: %s\n", config.LogFileName)
	} else {
		fmt.Println("Coral EAMS Client Started but log file name is empty")
	}

	logger.Info(fmt.Sprintf("Coral EAMS Client Started on %s:%d", config.HostName, config.Port))
}
