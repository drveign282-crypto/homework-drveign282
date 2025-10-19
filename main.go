package main

import "fmt"

func main() {
	ctrl()
}

// 控制流程-只出现一次的数字
func ctrl() {
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

// 控制流程-回文数：指正读和反读都相同的数字
// 只反转数字的后半部分，然后与前半部分比较，这样可以避免完全反转整个数字带来的溢出风险，尤其是大数
func isPalindrome(x int) bool {
	// 个位数和0直接返回false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 反转数初始化
	reversed := 0

	// 只需要反转一半的数字，和前半段对比就可以判断是否回文，如果是for x > 0就是完全反转
	for x > reversed {
		// *10进行数字左移，%10进行数据右取
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}
