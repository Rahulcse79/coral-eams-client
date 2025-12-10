package system

import (
	"os/exec"
	"runtime"
	"strings"
	"coral-eams-client/internal/logger"
)


type OSInfo struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Kernel       string `json:"kernel"`
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
	name := "Windows"
	version := "Unknown"
	kernel := "Unknown"

	out, err := exec.Command("cmd", "/C", "ver").Output()
	if err != nil {
		logger.Error("Failed to get Windows kernel version", "error", err)
	} else {
		kernel = strings.TrimSpace(string(out))
	}

	outVer, err := exec.Command("cmd", "/C", "wmic os get Version").Output()
	if err != nil {
		logger.Error("Failed to get Windows OS version", "error", err)
	} else {
		lines := strings.Split(string(outVer), "\n")
		if len(lines) > 1 {
			version = strings.TrimSpace(lines[1])
		}
	}

	return name, version, kernel
}

func getLinuxOS() (string, string, string) {
	name := "Linux"
	version := "Unknown"
	kernel := "Unknown"

	outName, err := exec.Command("cat", "/etc/os-release").Output()
	if err != nil {
		logger.Error("Failed to read /etc/os-release", "error", err)
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
	}

	outKernel, err := exec.Command("uname", "-r").Output()
	if err != nil {
		logger.Error("Failed to get Linux kernel version", "error", err)
	} else {
		kernel = strings.TrimSpace(string(outKernel))
	}

	return name, version, kernel
}

func getMacOS() (string, string, string) {
	name := "macOS"
	version := "Unknown"
	kernel := "Unknown"

	outVer, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		logger.Error("Failed to get macOS version", "error", err)
	} else {
		version = strings.TrimSpace(string(outVer))
	}

	outKernel, err := exec.Command("uname", "-r").Output()
	if err != nil {
		logger.Error("Failed to get macOS kernel version", "error", err)
	} else {
		kernel = strings.TrimSpace(string(outKernel))
	}

	return name, version, kernel
}
