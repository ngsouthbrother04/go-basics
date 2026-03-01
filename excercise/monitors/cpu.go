package monitors

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct {
}

func (cm *CPUMonitor) Name() string {
	return "CPU"
}

func (cm *CPUMonitor) Check(ctx context.Context) (string, bool) {
	cpuStat, err := cpu.PercentWithContext(ctx, 100*time.Millisecond, false)

	if err != nil {
		return fmt.Sprintf("[CPU] could not get CPU stats: %v\n", err), false
	}

	if len(cpuStat) == 0 {
		return "[CPU] no cpu stats available", false
	}

	return fmt.Sprintf("%.2f%%", cpuStat[0]), cpuStat[0] > 80
}
