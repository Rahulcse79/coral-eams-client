package sender

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"coral-eams-client/internal/logger"
)

type Sender struct {
	ServerURL string
	Token     string
	Client    *http.Client
}

func NewSender(url, token string) *Sender {
	return &Sender{
		ServerURL: url,
		Token:     token,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *Sender) Send(data interface{}) bool {
	payload, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal payload: " + err.Error())
		return false
	}

	req, err := http.NewRequest("POST", s.ServerURL, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error("Failed to create request: " + err.Error())
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	if s.Token != "" {
		req.Header.Set("Authorization", "Bearer "+s.Token)
	}

	logger.Info("Sending payload to server", "url", s.ServerURL)

	resp, err := s.Client.Do(req)
	if err != nil {
		logger.Error("Request failed: " + err.Error())
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		logger.Info("Payload successfully delivered", "status", resp.StatusCode)
		return true
	}

	logger.Warn("Server returned non-success status", "status", resp.StatusCode)
	return false
}

func CollectAndSendData(sysInfo interface{}, s *Sender) {
	logger.Info("Starting data collection")
	logger.Debug("Collected system info", "data", sysInfo)

	success := s.Send(sysInfo)
	if success {
		logger.Info("Data sent successfully")
	} else {
		logger.Warn("Failed to send data")
	}
}