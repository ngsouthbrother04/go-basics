package processors

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	"nnama.com/excercise/models"
)

const (
	monitorInterval           = 1 * time.Second
	defaultTopN               = 5
	resourceThreshold         = 1.0
	bytesPerMB        float64 = 1024.0 * 1024.0
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStat, m models.Monitor) {
	defer wg.Done()

	ticker := time.NewTicker(monitorInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			value, alert := m.Check(ctx)

			stat := models.SystemStat{
				Name: m.Name(),
				Stat: value,
			}

			select {
			case statCh <- stat:
			case <-ctx.Done():
				return
			}

			if alert {
				AlertLog(stat)
			}
		}
	}
}

func GetTopProcesses(ctx context.Context) string {
	vmStat, err := mem.VirtualMemory()

	if err != nil {
		return "[GetTopProcesses] could not get memory info"
	}

	processes, err := process.ProcessesWithContext(ctx)

	if err != nil {
		return fmt.Sprintf("[GetTopProcesses] could not receive process info: %v\n", err)
	}

	stats := collectProcessStats(ctx, processes, vmStat.Total)
	cpuList := sortByCPU(stats)
	memList := sortByMemory(stats)

	sort.Slice(cpuList, func(i, j int) bool {
		return cpuList[i].CPUPercent > cpuList[j].CPUPercent
	})

	sort.Slice(memList, func(i, j int) bool {
		return memList[i].MemPercent > memList[j].MemPercent
	})

	topCPU := topN(cpuList, defaultTopN)
	topMem := topN(memList, defaultTopN)

	if err := ExportToCSV(topCPU, topMem); err != nil {
		fmt.Printf("[ExportToCSV] %v\n", err)
	}

	return formatProcessReport(topCPU, topMem)
}

func collectProcessStats(ctx context.Context, processes []*process.Process, totalMemory uint64) []models.ProcessStat {
	var wg sync.WaitGroup
	procChan := make(chan models.ProcessStat, len(processes))

	for _, p := range processes {
		wg.Add(1)

		go func(p *process.Process) {
			defer wg.Done()

			stat, ok := buildProcessStat(ctx, p, totalMemory)
			if !ok {
				return
			}

			select {
			case procChan <- stat:
			case <-ctx.Done():
			}
		}(p)
	}

	go func() {
		wg.Wait()
		close(procChan)
	}()

	stats := make([]models.ProcessStat, 0, len(processes))
	for stat := range procChan {
		stats = append(stats, stat)
	}

	return stats
}

func buildProcessStat(ctx context.Context, p *process.Process, totalMemory uint64) (models.ProcessStat, bool) {
	select {
	case <-ctx.Done():
		return models.ProcessStat{}, false
	default:
	}

	name, err := p.NameWithContext(ctx)
	if err != nil {
		return models.ProcessStat{}, false
	}

	cpuPer, err := p.CPUPercentWithContext(ctx)
	if err != nil {
		return models.ProcessStat{}, false
	}

	memInfo, err := p.MemoryInfoWithContext(ctx)
	if err != nil {
		return models.ProcessStat{}, false
	}

	ramPercent := (float64(memInfo.RSS) / float64(totalMemory)) * 100
	if ramPercent < resourceThreshold || cpuPer < resourceThreshold {
		return models.ProcessStat{}, false
	}

	createTime, err := p.CreateTimeWithContext(ctx)
	if err != nil {
		return models.ProcessStat{}, false
	}

	runningTime := time.Since(time.Unix(createTime/1000, 0))

	return models.ProcessStat{
		PID:         p.Pid,
		Name:        name,
		CPUPercent:  cpuPer,
		Memory:      memInfo.RSS,
		MemPercent:  ramPercent,
		RunningTime: runningTime,
	}, true
}

func sortByCPU(stats []models.ProcessStat) []models.ProcessStat {
	filtered := make([]models.ProcessStat, 0, len(stats))
	for _, stat := range stats {
		if stat.CPUPercent >= resourceThreshold {
			filtered = append(filtered, stat)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].CPUPercent > filtered[j].CPUPercent
	})

	return filtered
}

func sortByMemory(stats []models.ProcessStat) []models.ProcessStat {
	filtered := make([]models.ProcessStat, 0, len(stats))
	for _, stat := range stats {
		if stat.MemPercent >= resourceThreshold {
			filtered = append(filtered, stat)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].MemPercent > filtered[j].MemPercent
	})

	return filtered
}

func topN(list []models.ProcessStat, n int) []models.ProcessStat {
	if len(list) < n {
		return list
	}

	return list[:n]
}

func formatProcessReport(cpuList, memList []models.ProcessStat) string {
	var builder strings.Builder

	builder.WriteString("== Top 5 CPU cosuming processes == \n")
	for idx, v := range cpuList {
		builder.WriteString(fmt.Sprintf("%d. [%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running: %s \n",
			idx+1,
			v.PID,
			v.Name,
			v.CPUPercent,
			float64(v.Memory)/bytesPerMB,
			v.MemPercent,
			v.RunningTime,
		))
	}

	builder.WriteString("== Top 5 RAM cosuming processes == \n")
	for idx, v := range memList {
		builder.WriteString(fmt.Sprintf("%d. [%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running: %s \n",
			idx+1,
			v.PID,
			v.Name,
			v.CPUPercent,
			float64(v.Memory)/bytesPerMB,
			v.MemPercent,
			v.RunningTime,
		))
	}

	return builder.String()
}

func ExportToCSV(cpuList, memList []models.ProcessStat) error {
	file, err := os.OpenFile("process_stats.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		if _, err := file.WriteString("Timestamp,PID,Name,CPU (%),RAM (MB),RAM (%),Running Time \n"); err != nil {
			return fmt.Errorf("could not write csv header: %w", err)
		}
	}

	timestamp := time.Now().Format(time.RFC3339)
	for _, c := range cpuList {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s\n",
			timestamp,
			c.PID,
			c.Name,
			c.CPUPercent,
			float64(c.Memory)/bytesPerMB,
			c.MemPercent,
			c.RunningTime,
		)
		if _, err := file.WriteString(line); err != nil {
			return fmt.Errorf("could not write cpu row: %w", err)
		}
	}

	for _, c := range memList {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s\n",
			timestamp,
			c.PID,
			c.Name,
			c.CPUPercent,
			float64(c.Memory)/bytesPerMB,
			c.MemPercent,
			c.RunningTime,
		)
		if _, err := file.WriteString(line); err != nil {
			return fmt.Errorf("could not write memory row: %w", err)
		}
	}

	return nil
}

func AlertLog(stat models.SystemStat) {
	models.Mutex.Lock()
	defer models.Mutex.Unlock()

	file, err := os.OpenFile("alert.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("[AlertLog] could not open file: %v\n", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format(time.RFC3339)
	logLine := fmt.Sprintf("[%s] ALERT: %s = %s \n", timestamp, stat.Name, stat.Stat)
	if _, err := file.WriteString(logLine); err != nil {
		fmt.Printf("[AlertLog] could not write log line: %v\n", err)
	}
}
