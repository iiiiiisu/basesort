package main

import (
	"basesort/methods"
	"basesort/utils"
	"basesort/utils/log"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	logger := log.ConsoleLog{}
	logger.Init()
	m := utils.SortManager{
		Length: 10000,
		Seed:   0,
		Times: 1,
		Log: &logger,
		Sorts: []methods.Sorter{
			&methods.InsertionSort{},
			&methods.SelectionSort{},
			&methods.BubbleSort{},
			&methods.ShellSort{},
			&methods.QuickSort{},
			&methods.MergeSort{},
			&methods.HeapSort{},
		},
	}
	for i := 0; i < m.Times; i ++ {
		m.Start()
		m.Seed = time.Now().UnixNano()
	}
}
