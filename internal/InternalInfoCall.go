package system

import (
	"coral-eams-client/internal/logger"
)

func InternalInfoCall() *SystemInfo {
	
	sysInfo := CollectSystemInfo()
	logger.Info("Internal system info collected via InternalInfoCall")
	return sysInfo
}
