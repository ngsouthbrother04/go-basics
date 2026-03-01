package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

type NetMonitor struct {
}

func (cm *NetMonitor) Name() string {
	return "Net"
}

func (cm *NetMonitor) Check(ctx context.Context) (string, bool) {
	netStat, err := net.IOCountersWithContext(ctx, false)

	if err != nil {
		return fmt.Sprintf("[Net Check] could not get net info: %v\n", err), false
	}

	if len(netStat) == 0 {
		return "[Net Check] no net stats available", false
	}

	return fmt.Sprintf("Sent: %d KB, Recv: %d KB", netStat[0].BytesSent/1024, netStat[0].BytesRecv/1024), false
}
