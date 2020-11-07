package main

import (
	"basesort/methods"
	"basesort/methods/tree"
	"basesort/utils"
	"basesort/utils/log"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	logger := log.FileLog{}
	logger.Init("log.csv")
	defer logger.Close()
	defer logger.Output()
	m := utils.SortManager{
		Length: 100000,
		Seed:   0,
		Times:  3,
		Log:    &logger,
		Sorts: []methods.Sorter{
			&methods.InsertionSort{},
			&methods.SelectionSort{},
			&methods.BubbleSort{},
			&methods.ShellSort{},
			&methods.QuickSort{},
			&methods.MergeSort{},
			&methods.HeapSort{},
			&tree.BalanceBinary{},
		},
	}
	for i := 0; i < m.Times; i++ {
		m.Start()
		m.Seed = time.Now().UnixNano()
	}
}
