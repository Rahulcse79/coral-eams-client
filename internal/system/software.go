package system

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
	"coral-eams-client/internal/logger"
)

func GetInstalledSoftware() []string {
	var softwareList []string

	switch runtime.GOOS {
	case "windows":
		softwareList = getWindowsSoftware()
	case "linux":
		softwareList = getLinuxSoftware()
	case "darwin":
		softwareList = getMacSoftware()
	default:
		logger.Warn("Unsupported OS for software collection")
	}

	if len(softwareList) > 0 {
		logger.Info("Installed software info collected", "softwareList", softwareList)
	} else {
		logger.Info("No installed software found")
	}

	return softwareList
}

func getWindowsSoftware() []string {
	var list []string
	out, err := exec.Command("powershell", "-Command",
		`Get-ItemProperty HKLM:\Software\Microsoft\Windows\CurrentVersion\Uninstall\* | Select-Object DisplayName`).Output()
	if err != nil {
		logger.Error(err)
		return list
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && line != "DisplayName" {
			list = append(list, line)
		}
	}
	return list
}

func getLinuxSoftware() []string {
	var list []string

	out, err := exec.Command("bash", "-c", "dpkg -l 2>/dev/null || rpm -qa 2>/dev/null").Output()
	if err != nil {
		logger.Error(err)
		return list
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			list = append(list, line)
		}
	}
	return list
}

func getMacSoftware() []string {
	var list []string
	out, err := exec.Command("bash", "-c", "system_profiler SPApplicationsDataType | grep 'Location'").Output()
	if err != nil {
		logger.Error(err)
		return list
	}

	lines := bytes.Split(out, []byte("\n"))
	for _, line := range lines {
		text := strings.TrimSpace(string(line))
		if text != "" {
			list = append(list, text)
		}
	}
	return list
}
