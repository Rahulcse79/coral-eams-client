package system

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"

	"coral-eams-client/internal/logger"
)

type ServiceInfo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// GetServicesInfo collects running services for the current OS
func GetServicesInfo() []ServiceInfo {
	var services []ServiceInfo

	switch runtime.GOOS {
	case "windows":
		services = getWindowsServices()
	case "linux":
		services = getLinuxServices()
	case "darwin":
		services = getMacServices()
	default:
		logger.Warn("Unsupported OS for fetching services")
	}

	logger.Info("Services info collected", "totalServices", len(services))
	return services
}

func getWindowsServices() []ServiceInfo {
	var services []ServiceInfo
	out, err := exec.Command("powershell", "-Command", "Get-Service | Select-Object Name, Status").Output()
	if err != nil {
		logger.Error("Failed to fetch Windows services", "error", err)
		return services
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "Name") || strings.Contains(line, "----") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			services = append(services, ServiceInfo{Name: fields[0], Status: fields[1]})
		}
	}
	return services
}

func getLinuxServices() []ServiceInfo {
	var services []ServiceInfo
	out, err := exec.Command("systemctl", "list-units", "--type=service", "--state=running").Output()
	if err != nil {
		logger.Error("Failed to fetch Linux services", "error", err)
		return services
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			services = append(services, ServiceInfo{Name: fields[0], Status: fields[3]})
		}
	}
	return services
}

func getMacServices() []ServiceInfo {
	var services []ServiceInfo
	out, err := exec.Command("launchctl", "list").Output()
	if err != nil {
		logger.Error("Failed to fetch macOS services", "error", err)
		return services
	}
	lines := bytes.Split(out, []byte("\n"))
	for _, line := range lines {
		fields := bytes.Fields(line)
		if len(fields) >= 3 {
			status := string(fields[0])
			if status == "0" {
				status = "running"
			} else {
				status = "stopped"
			}
			services = append(services, ServiceInfo{Name: string(fields[2]), Status: status})
		}
	}
	return services
}
