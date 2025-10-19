package main

import "fmt"

func main() {
	m := make(map[int]int)
	nums := [5]int{4, 1, 2, 1, 2}
	for _, num := range nums {
		// 如果索引不存在，则返回0
		if m[num] == 0 {
			m[num] = 1
		} else {
			// value自增，简写
			m[num]++
		}
	}
	// 获取value的值为1的元素
	for k, v := range m {
		if v == 1 {
			fmt.Println(k)
		}
	}
}
