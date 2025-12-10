package system

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
	"strings"
	"coral-eams-client/internal/logger"
)

type OSInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Kernel      string `json:"kernel"`
	Architecture string `json:"architecture"`
}

func GetOSInfo() *OSInfo {
	osInfo := &OSInfo{}
	osInfo.Architecture = runtime.GOARCH

	switch runtime.GOOS {
	case "windows":
		osInfo.Name, osInfo.Version, osInfo.Kernel = getWindowsOS()
	case "linux":
		osInfo.Name, osInfo.Version, osInfo.Kernel = getLinuxOS()
	case "darwin":
		osInfo.Name, osInfo.Version, osInfo.Kernel = getMacOS()
	default:
		osInfo.Name = "Unknown"
		osInfo.Version = "Unknown"
		osInfo.Kernel = "Unknown"
		logger.Warn("Unsupported OS detected")
	}

	logger.Info("OS info collected",
		"Name", osInfo.Name,
		"Version", osInfo.Version,
		"Kernel", osInfo.Kernel,
		"Architecture", osInfo.Architecture,
	)

	return osInfo
}

func getWindowsOS() (string, string, string) {
	name := ""
	version := ""
	kernel := ""

	out, err := exec.Command("cmd", "/C", "ver").Output()
	if err != nil {
		logger.Error(err)
		name = "Windows"
		version = "Unknown"
		kernel = "Unknown"
		return name, version, kernel
	}

	kernel = strings.TrimSpace(string(out))
	name = "Windows"

	outVer, err := exec.Command("cmd", "/C", "wmic os get Version").Output()
	if err != nil {
		logger.Error(err)
		version = "Unknown"
	} else {
		lines := strings.Split(string(outVer), "\n")
		if len(lines) > 1 {
			version = strings.TrimSpace(lines[1])
		} else {
			version = "Unknown"
		}
	}

	return name, version, kernel
}

func getLinuxOS() (string, string, string) {
	name := ""
	version := ""
	kernel := ""

	outName, err := exec.Command("cat", "/etc/os-release").Output()
	if err != nil {
		logger.Error(err)
		name = "Linux"
		version = "Unknown"
	} else {
		lines := strings.Split(string(outName), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				name = strings.Trim(line[len("PRETTY_NAME="):], `"`)
			}
			if strings.HasPrefix(line, "VERSION_ID=") {
				version = strings.Trim(line[len("VERSION_ID="):], `"`)
			}
		}
		if name == "" {
			name = "Linux"
		}
		if version == "" {
			version = "Unknown"
		}
	}

	outKernel, err := exec.Command("uname", "-r").Output()
	if err != nil {
		logger.Error(err)
		kernel = "Unknown"
	} else {
		kernel = strings.TrimSpace(string(outKernel))
	}

	return name, version, kernel
}

func getMacOS() (string, string, string) {
	name := "macOS"
	version := ""
	kernel := ""

	outVer, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		logger.Error(err)
		version = "Unknown"
	} else {
		version = strings.TrimSpace(string(outVer))
	}

	outKernel, err := exec.Command("uname", "-r").Output()
	if err != nil {
		logger.Error(err)
		kernel = "Unknown"
	} else {
		kernel = strings.TrimSpace(string(outKernel))
	}

	return name, version, kernel
}
