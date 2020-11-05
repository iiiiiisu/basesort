# basesort
数据结构与算法里面几种基础的排序算法。Golang实现 （数组）

包括：
- 插入 (InsertionSort)
- 选择 (SelectionSort)
- 冒泡 (BubbleSort)
- 希尔 (ShellSort)
- 快速 (QuickSort)
- 归并 (MergeSort)
- 堆 (HeapSort)

每个排序都有_test.go (测试) go test

每个算法都未进行简单逻辑优化。

util包中提供了SortManager结构体, 和log结构体，可以通过编写log结构体实现结果的输出方式，目前简单提供了一个控制台即时输出。
```
// 初始化一个log，选择测试结果输出方式
logger := log.ConsoleLog{}
logger.Init()
// 初始化一个SortManager
// Length ：测试数据长度
// seed ： 随机种子，为0时不设置随机种子
// Times： 执行次数，并未在代码中实现多次执行，需要手动循环 Times次
// Log: logger接口，使用不同的结构体可以有不同的输出方式
// Sorts：测试哪些排序方法，（注册排序方法）。
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
  // Start() 执行顺序：
  // 新建一个长度为Length的数组并保存，
  // 打乱保存的数组
  // 遍历 Sorts，Copy已打乱的数组，执行每个排序方法的排序：
  m.Start()
  m.Seed = time.Now().UnixNano()
}
```

> 欢迎看到的各位提意见，收到会及时采纳。
