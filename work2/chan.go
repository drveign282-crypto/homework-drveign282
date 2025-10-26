package work2

import (
	"fmt"
	"sync"
)

func Sender(c chan<- int, wg *sync.WaitGroup, num int) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		c <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(c)
}

func Receive(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range c {
		fmt.Printf("接收: %d\n", i)
	}
}
