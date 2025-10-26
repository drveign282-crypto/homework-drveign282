package work2

import (
	"fmt"
	"sync"
	"time"
)

// 打印从0到1的奇数
func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			println(i)
		}
	}
}

// 打印从0到1的偶数
func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}
}

func SyncWaitGroupTest() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go printOdd(&wg)
	go printEven(&wg)
	wg.Wait()
	println("done")
}

func TaskExecute(tasks []func()) {
	var wg = sync.WaitGroup{}
	wg.Add(len(tasks))
	// 创建多个协程执行任务
	for i, task := range tasks {
		// 执行任务，并记录任务执行时间。
		go func(id int, f func()) {
			defer wg.Done()
			start := time.Now()
			f()
			elapsed := time.Since(start)
			fmt.Printf("任务%d 执行耗时: %v\n", id, elapsed)
		}(i, task)
	}
	wg.Wait()
	println("done")
}
