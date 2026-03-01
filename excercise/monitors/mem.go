package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryMonitor struct{}

func (rm *MemoryMonitor) Name() string {
	return "Memory"
}

func (rm *MemoryMonitor) Check(ctx context.Context) (string, bool) {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Memory Check] could not get memory info: %v\n", err), false
	}

	return fmt.Sprintf("%.2f%%", vmStat.UsedPercent), vmStat.UsedPercent > 80
}
