package log

import "fmt"

type ConsoleLog struct {
	Logs []SortLog
}

func (l *ConsoleLog) Init() {
	l.Logs = make([]SortLog, 1)
}

func (l *ConsoleLog) Add(log SortLog) {
	l.Logs[0] = log
	l.Output()
}

func (l *ConsoleLog) Output() {
	log := l.Logs[0]
	fmt.Printf("%-16s %7d %s %10s %12d %#v %d\n",
		log.Name, log.Length, log.Start.Format("2006-01-02 15:04:05"),
		log.TimeCount, int64(log.TimeCount), log.Result, log.Seed)
}
