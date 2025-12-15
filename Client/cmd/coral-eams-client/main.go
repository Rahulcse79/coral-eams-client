package main

import (
	"coral-eams-client/internal/logger"
	"coral-eams-client/internal/scheduler"
	"coral-eams-client/internal/sender"
	"coral-eams-client/internal/system"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	LogFileName           string `json:"logFileName"`
	Port                  int    `json:"port"`
	HostName              string `json:"hostName"`
	SchedulerTimeInterval int    `json:"schedulerTimeInterval"`
	SenderURL             string `json:"SenderURL"`
	AuthToken             string `json:"AuthToken"`
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

	if config.LogFileName != "" {
		fmt.Printf("Coral EAMS Client Started and log file created: %s\n", config.LogFileName)
		logger.InitLogger(config.LogFileName)
	} else {
		fmt.Println("Coral EAMS Client Started but log file name is empty")
		logger.InitLogger("coral-eams-client-log.log")
	}

	logger.Info(fmt.Sprintf("Coral EAMS Client Started on %s:%d", config.HostName, config.Port))

	sysInfo := system.CollectSystemInfo()
	logger.Info(fmt.Sprintf("Collected system info: %+v\n", sysInfo))

	macAddress := system.CollectMacAddressInfo()
	firstMAC := "none"
	if macAddress != nil && len(macAddress.MacAddress) > 0 && len(macAddress.MacAddress[0]) > 0 {
		firstMAC = macAddress.MacAddress[0]
	}
	logger.Info(fmt.Sprintf("Collected macAddress info: %s", firstMAC))

	logger.Info("Sender initialized", "serverURL", config.SenderURL+"/"+firstMAC)
	if config.SenderURL != "" {
		s := sender.NewSender(config.SenderURL+"/"+firstMAC, config.AuthToken)
		interval := config.SchedulerTimeInterval
		if interval <= 0 {
			interval = 5
			logger.Warn("Invalid schedulerTimeInterval — using default 5 minutes")
		}

		logger.Info("Scheduler initialized", "intervalMinutes", config.SchedulerTimeInterval)
		cronJob := scheduler.NewCronJob(interval)
		cronJob.Start(func() {
			logger.Info("Starting scheduled data send")
			sender.CollectAndSendData(sysInfo, s)
		})

		select {}

	} else {
		logger.Warn("Sender URL is empty in configuration")
	}
}
