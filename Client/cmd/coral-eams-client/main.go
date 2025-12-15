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

	logger.Info(fmt.Sprintf(
		"Coral EAMS Client Started on %s:%d",
		config.HostName,
		config.Port,
	))

	sysInfo := system.CollectSystemInfo()
	logger.Info("Collected system info")

	macInfo := system.CollectMacAddressInfo()
	firstMAC := "none"

	if macInfo != nil &&
		len(macInfo.MacAddress) > 0 &&
		len(macInfo.MacAddress[0]) > 0 {
		firstMAC = macInfo.MacAddress[0]
	}

	logger.Info("Collected macAddress info", "mac", firstMAC)

	if config.SenderURL == "" {
		logger.Warn("Sender URL is empty in configuration")
		select {}
	}

	fullURL := config.SenderURL + "/" + firstMAC
	logger.Info("Sender initialized", "serverURL", fullURL)

	s := sender.NewSender(fullURL, config.AuthToken)

	interval := config.SchedulerTimeInterval
	if interval <= 0 {
		interval = 1440
		logger.Warn("Invalid schedulerTimeInterval — using default one day interval")
	}

	logger.Info(
		"Scheduler initialized",
		"intervalMinutes", interval,
		"intervalHours", interval/60,
	)

	logger.Info("Sending system info immediately on startup")
	sender.CollectAndSendData(sysInfo, s)

	cronJob := scheduler.NewCronJob(interval)
	cronJob.Start(func() {
		logger.Info("Starting scheduled data send")
		sender.CollectAndSendData(sysInfo, s)
	})

	select {}
}
