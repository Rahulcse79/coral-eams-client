package system

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"

	"coral-eams-client/internal/logger"
)

type HardwareInfo struct {
	CPUModel     string  `json:"cpu_model"`
	CPUCores     int     `json:"cpu_cores"`
	RAMGB        float64 `json:"ram_gb"`
	DiskTotalGB  float64 `json:"disk_total_gb"`
	SerialNumber string  `json:"serial_number"`
	Motherboard  string  `json:"motherboard"`
}

func GetHardwareInfo() *HardwareInfo {
	hw := &HardwareInfo{}

	// CPU info
	cpuInfo, err := cpu.Info()
	if err != nil || len(cpuInfo) == 0 {
		logger.Error("Failed to get CPU info", "error", err)
	} else {
		hw.CPUModel = cpuInfo[0].ModelName
		hw.CPUCores = int(cpuInfo[0].Cores)
	}

	// RAM info
	vm, err := mem.VirtualMemory()
	if err != nil {
		logger.Error("Failed to get RAM info", "error", err)
	} else {
		hw.RAMGB = float64(vm.Total) / (1024 * 1024 * 1024)
	}

	// Disk info
	diskInfo, err := disk.Usage("/")
	if err != nil {
		logger.Error("Failed to get Disk info", "error", err)
	} else {
		hw.DiskTotalGB = float64(diskInfo.Total) / (1024 * 1024 * 1024)
	}

	// Serial & Motherboard
	switch runtime.GOOS {
	case "windows":
		hw.SerialNumber, _ = getWindowsSerial()
		hw.Motherboard, _ = getWindowsMotherboard()
	case "linux":
		hw.SerialNumber, _ = getLinuxSerial()
		hw.Motherboard, _ = getLinuxMotherboard()
	case "darwin":
		hw.SerialNumber, _ = getMacSerial()
		hw.Motherboard = "Apple Logic Board"
	default:
		hw.SerialNumber = "Unknown"
		hw.Motherboard = "Unknown"
		logger.Warn("Unsupported OS for hardware info collection")
	}

	logger.Info("Hardware info collected",
		"CPUModel", hw.CPUModel,
		"CPUCores", hw.CPUCores,
		"RAMGB", hw.RAMGB,
		"DiskTotalGB", hw.DiskTotalGB,
		"SerialNumber", hw.SerialNumber,
		"Motherboard", hw.Motherboard,
	)

	return hw
}

// ---------------- Windows ----------------

func getWindowsSerial() (string, error) {
	out, err := exec.Command("wmic", "bios", "get", "serialnumber").Output()
	if err != nil {
		logger.Error("Failed to get Windows serial", "error", err)
		return "", err
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		err := errors.New("serial number not found")
		logger.Error("Windows serial number not found", "error", err)
		return "", err
	}
	return strings.TrimSpace(lines[1]), nil
}

func getWindowsMotherboard() (string, error) {
	out, err := exec.Command("wmic", "baseboard", "get", "product").Output()
	if err != nil {
		logger.Error("Failed to get Windows motherboard", "error", err)
		return "", err
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		err := errors.New("motherboard info not found")
		logger.Error("Windows motherboard info not found", "error", err)
		return "", err
	}
	return strings.TrimSpace(lines[1]), nil
}

// ---------------- Linux ----------------

func getLinuxSerial() (string, error) {
	out, err := exec.Command("cat", "/sys/class/dmi/id/product_serial").Output()
	if err != nil {
		logger.Error("Failed to get Linux serial", "error", err)
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getLinuxMotherboard() (string, error) {
	out, err := exec.Command("cat", "/sys/class/dmi/id/board_name").Output()
	if err != nil {
		logger.Error("Failed to get Linux motherboard", "error", err)
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// ---------------- Mac ----------------

func getMacSerial() (string, error) {
	out, err := exec.Command("system_profiler", "SPHardwareDataType").Output()
	if err != nil {
		logger.Error("Failed to get Mac serial", "error", err)
		return "", err
	}
	lines := bytes.Split(out, []byte("\n"))
	for _, line := range lines {
		if bytes.Contains(line, []byte("Serial Number")) {
			parts := bytes.Split(line, []byte(":"))
			if len(parts) > 1 {
				return strings.TrimSpace(string(parts[1])), nil
			}
		}
	}
	err = errors.New("serial number not found")
	logger.Error("Mac serial number not found", "error", err)
	return "", err
}
