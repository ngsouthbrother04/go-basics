package monitors

import (
	"context"
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskMonitor struct {
}

func (cm *DiskMonitor) Name() string {
	return "Disk"
}

func (cm *DiskMonitor) Check(ctx context.Context) (string, bool) {
	path := "/"

	if runtime.GOOS == "windows" {
		path = "C:"
	}

	diskStat, err := disk.UsageWithContext(ctx, path)

	if err != nil {
		return fmt.Sprintf("[Disk Check] could not get disk stats: %v\n", err), false
	}

	return fmt.Sprintf("Used: %.2f%%, Free: %d, Total: %d", diskStat.UsedPercent, diskStat.Free, diskStat.Total), diskStat.UsedPercent > 80
}
