package log

import "time"

type Logger interface {
	Add(SortLog)
	Output()
}

type SortLog struct {
	Name string
	Seed int64
	Length int
	Start time.Time
	TimeCount time.Duration
	Result bool
}