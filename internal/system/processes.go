package system

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"coral-eams-client/internal/logger"
)

type ProcessInfo struct {
	PID        int32   `json:"pid"`
	Name       string  `json:"name"`
	CPUPercent float64 `json:"cpu_percent"`
	MemPercent float32 `json:"mem_percent"`
}

func GetProcessesInfo() []ProcessInfo {
	var processesInfo []ProcessInfo

	procs, err := process.Processes()
	if err != nil {
		logger.Error("Failed to fetch processes info", "error", err)
		return processesInfo
	}

	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			logger.Debug(fmt.Sprintf("Failed to get name for PID %d: %v", p.Pid, err))
			continue
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			logger.Debug(fmt.Sprintf("Failed to get CPU percent for PID %d: %v", p.Pid, err))
			continue
		}

		memPercent, err := p.MemoryPercent()
		if err != nil {
			logger.Debug(fmt.Sprintf("Failed to get Memory percent for PID %d: %v", p.Pid, err))
			continue
		}

		processesInfo = append(processesInfo, ProcessInfo{
			PID:        p.Pid,
			Name:       name,
			CPUPercent: cpuPercent,
			MemPercent: memPercent,
		})
	}

	logger.Info(fmt.Sprintf("Collected %d processes info", len(processesInfo)))
	
	return processesInfo
}
