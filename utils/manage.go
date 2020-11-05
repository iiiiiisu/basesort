package utils

import (
	"basesort/methods"
	"basesort/utils/log"
	"fmt"
	"reflect"
	"strings"
	"time"
)



type SortManager struct {
	Sorts []methods.Sorter
	Length int
	Seed int64
	Times int
	Log log.Logger
}

func (m *SortManager) Start(){
	data := NewArray(m.Length)
	if m.Seed == 0 {
		Shuffle(data)
	} else {
		Shuffle(data, m.Seed)
	}
	bak := make([]uint32, m.Length)
	for _, s := range m.Sorts {
		copy(bak, data)
		s.SetData(bak)
		startTime := time.Now()
		s.Sort()
		elapsedTime := time.Since(startTime)
		result := s.CheckResult()
		name := fmt.Sprintf("%s", reflect.TypeOf(s))
		m.Log.Add(log.SortLog{
			Name: strings.TrimLeft(name, "*methods."),
			Seed: m.Seed,
			Length: m.Length,
			Start: startTime,
			TimeCount: elapsedTime,
			Result: result,
		})
	}
}
