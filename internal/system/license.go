package system

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"coral-eams-client/internal/logger"
)

// LicenseInfo holds the product key and OS identifier
type LicenseInfo struct {
	ProductKey   string `json:"product_key"`
	OSIdentifier string `json:"os_identifier"`
}

// GetLicenseInfo collects license info based on OS
func GetLicenseInfo() *LicenseInfo {
	license := &LicenseInfo{}

	switch runtime.GOOS {
	case "windows":
		key, err := getWindowsLicenseKey()
		if err != nil {
			key = "Unknown"
		}
		license.ProductKey = key

		osid, err := getWindowsOSID()
		if err != nil {
			osid = "Unknown"
		}
		license.OSIdentifier = osid

	case "linux":
		key, err := getLinuxMachineID()
		if err != nil {
			key = "Unknown"
		}
		license.ProductKey = key
		license.OSIdentifier = "Linux"

	case "darwin":
		key, err := getMacSerial() // reuse hardware.go
		if err != nil {
			key = "Unknown"
		}
		license.ProductKey = key
		license.OSIdentifier = "macOS"

	default:
		license.ProductKey = "Unknown"
		license.OSIdentifier = "Unknown"
	}

	logger.Info("License info collected",
		"ProductKey", license.ProductKey,
		"OSIdentifier", license.OSIdentifier,
	)

	return license
}

// ---------------- Windows ----------------

func getWindowsLicenseKey() (string, error) {
	out, err := exec.Command("powershell", "-Command",
		"(Get-WmiObject -query 'select * from SoftwareLicensingService').OA3xOriginalProductKey").Output()
	if err != nil {
		logger.Error("Failed to get Windows license key", "error", err)
		return "", err
	}
	key := strings.TrimSpace(string(out))
	if key == "" {
		err = errors.New("windows license key not found")
		logger.Error("Windows license key not found", "error", err)
		return "", err
	}
	return key, nil
}

func getWindowsOSID() (string, error) {
	out, err := exec.Command("wmic", "os", "get", "SerialNumber").Output()
	if err != nil {
		logger.Error("Failed to get Windows OS ID", "error", err)
		return "", err
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.EqualFold(line, "SerialNumber") {
			return line, nil
		}
	}
	err = errors.New("windows OS ID not found")
	logger.Error("Windows OS ID not found", "error", err)
	return "", err
}

// ---------------- Linux ----------------

func getLinuxMachineID() (string, error) {
	if _, err := os.Stat("/etc/machine-id"); err != nil {
		logger.Error("Linux machine-id file missing", "error", err)
		return "", err
	}
	out, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		logger.Error("Failed to read Linux machine-id", "error", err)
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
