package system

import (
	"coral-eams-client/internal/logger"
)

type SystemInfo struct {
	Hardware  *HardwareInfo  `json:"hardware"`
	OS        *OSInfo        `json:"os"`
	Network   []*NetworkInfo `json:"network"`
	Software  []string       `json:"software"`
	License   *LicenseInfo   `json:"license"`
	Processes []ProcessInfo  `json:"processes"`
	Services  []ServiceInfo  `json:"services"`
}

type MacAddress struct {
	MacAddress []string `json:"MacAddress"`
}

func CollectMacAddressInfo() *MacAddress {

	macAddress := &MacAddress{}
	macAddress.MacAddress = GetAllMACAddresses()

	return macAddress
}

func CollectSystemInfo() *SystemInfo {

	info := &SystemInfo{}

	info.Hardware = GetHardwareInfo()
	logger.Info("Hardware info collected",
		"CPUModel", info.Hardware.CPUModel,
		"CPUCores", info.Hardware.CPUCores,
	)

	info.OS = GetOSInfo()
	logger.Info("OS info collected",
		"OSName", info.OS.Name,
		"OSVersion", info.OS.Version,
	)

	info.Network = GetNetworkInfo()
	logger.Info("Network info collected",
		"InterfacesCount", len(info.Network),
	)

	info.Software = GetInstalledSoftware()
	logger.Info("Installed software info collected",
		"SoftwareCount", len(info.Software),
	)

	info.License = GetLicenseInfo()
	logger.Info("License info collected",
		"ProductKey", info.License.ProductKey,
		"OSIdentifier", info.License.OSIdentifier,
	)

	info.Processes = GetProcessesInfo()
	logger.Info("Processes info collected",
		"ProcessesCount", len(info.Processes),
	)

	info.Services = GetServicesInfo()
	logger.Info("Services info collected",
		"ServicesCount", len(info.Services),
	)

	return info
}
