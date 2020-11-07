package log

import (
	"fmt"
	"os"
)

type FileLog struct {
	File *os.File
	Logs []SortLog
}

func (l *FileLog) Init(file string) error {
	var err error
	if l.File, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend|os.ModePerm); err != nil {
		return err
	}
	return nil
}
func (l *FileLog) Add(log SortLog) {
	l.Logs = append(l.Logs, log)
	if len(l.Logs) > 10 {
		l.Output()
	}
}

func (l *FileLog) Output() {
	if len(l.Logs) == 0 {
		return
	}
	f := l.File
	for _, log := range l.Logs {
		text := fmt.Sprintf("%s,%d,%s,%s,%d,%#v,%d\n",
			log.Name, log.Length, log.Start.Format("2006-01-02 15:04:05"),
			log.TimeCount, int64(log.TimeCount), log.Result, log.Seed)
		f.WriteString(text)
	}
	l.Logs = []SortLog{}
}

func (l *FileLog) Close() {
	l.File.Close()
}
