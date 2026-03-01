package models

import (
	"context"
	"sync"
	"time"
)

type SystemStat struct {
	Name string
	Stat string
}

type Monitor interface {
	Name() string
	Check(ctx context.Context) (string, bool)
}

type ProcessStat struct {
	PID         int32
	Name        string
	CPUPercent  float64
	Memory      uint64
	MemPercent  float64
	RunningTime time.Duration
}

type StatsStore struct {
	mu    sync.RWMutex
	stats map[string]SystemStat
}

func NewStatsStore() *StatsStore {
	return &StatsStore{
		stats: make(map[string]SystemStat),
	}
}

func (s *StatsStore) Set(stat SystemStat) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stats[stat.Name] = stat
}

func (s *StatsStore) Snapshot() map[string]SystemStat {
	s.mu.RLock()
	defer s.mu.RUnlock()

	clone := make(map[string]SystemStat, len(s.stats))
	for key, value := range s.stats {
		clone[key] = value
	}

	return clone
}

var (
	Stats = NewStatsStore()
	Mutex sync.Mutex
)
