package system

import (
	"net"
	"runtime"
	"strings"

	"coral-eams-client/internal/logger"
)

type NetworkInfo struct {
	InterfaceName string `json:"interface_name"`
	MAC           string `json:"mac"`
	IPv4          string `json:"ipv4"`
	IPv6          string `json:"ipv6"`
	IsUp          bool   `json:"is_up"`
}

// GetNetworkInfo returns info about all active network interfaces
func GetNetworkInfo() []*NetworkInfo {
	var result []*NetworkInfo
	defer logger.Info("Network info collection completed", "interfaces_count", len(result))

	if runtime.GOOS != "windows" && runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		logger.Warn("Unsupported OS for network info")
		return result
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		logger.Error("Failed to get network interfaces", "error", err)
		return result
	}

	for _, iface := range interfaces {
		// Skip down interfaces
		isUp := iface.Flags&net.FlagUp != 0
		ifaceInfo := &NetworkInfo{
			InterfaceName: iface.Name,
			MAC:           iface.HardwareAddr.String(),
			IsUp:          isUp,
		}

		addrs, err := iface.Addrs()
		if err != nil {
			logger.Error("Failed to get interface addresses", "interface", iface.Name, "error", err)
			continue
		}

		for _, addr := range addrs {
			ip := strings.Split(addr.String(), "/")[0] // remove subnet
			if strings.Contains(ip, ":") {
				ifaceInfo.IPv6 = ip
			} else {
				ifaceInfo.IPv4 = ip
			}
		}

		result = append(result, ifaceInfo)
	}

	return result
}

// GetPrimaryIP returns the first non-loopback IPv4 address found
func GetPrimaryIP() string {
	defer logger.Info("Primary IP fetched")

	interfaces, err := net.Interfaces()
	if err != nil {
		logger.Error("Failed to get network interfaces", "error", err)
		return ""
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ip := strings.Split(addr.String(), "/")[0]
			if strings.Contains(ip, ".") && !strings.HasPrefix(ip, "127.") {
				return ip
			}
		}
	}

	return ""
}

// GetPrimaryMAC returns the MAC of the first active interface
func GetPrimaryMAC() string {
	defer logger.Info("Primary MAC fetched")

	interfaces, err := net.Interfaces()
	if err != nil {
		logger.Error("Failed to get network interfaces", "error", err)
		return ""
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String()
		}
	}

	return ""
}
