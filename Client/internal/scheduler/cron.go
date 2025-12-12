package scheduler

import (
	"time"
	"coral-eams-client/internal/logger"
)

type CronJob struct {
	IntervalMinutes int
	stopChan        chan bool
	running         bool
}

func NewCronJob(interval int) *CronJob {
	return &CronJob{
		IntervalMinutes: interval,
		stopChan:        make(chan bool),
	}
}

func (c *CronJob) Start(task func()) {
	if c.running {
		logger.Warn("Cron job already running")
		return
	}
	c.running = true

	go func() {
		logger.Info("Cron job started", "intervalMinutes", c.IntervalMinutes)

		ticker := time.NewTicker(time.Duration(c.IntervalMinutes) * time.Minute)

		for {
			select {
			case <-ticker.C:
				logger.Debug("Cron tick triggered")
				task()

			case <-c.stopChan:
				ticker.Stop()
				c.running = false
				logger.Info("Cron job stopped")
				return
			}
		}
	}()
}

func (c *CronJob) Stop() {
	if !c.running {
		logger.Warn("Cron job stop requested but it is not running")
		return
	}
	c.stopChan <- true
}
