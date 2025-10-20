package main

import (
	"fmt"
	"sort"
)

func main() {
	ctrl()
	fmt.Println(isPalindrome(12321))
	fmt.Println(isValid("[}()"))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(plusOne([]int{1, 2, 9}))
	fmt.Print(rmDuplicates([]int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 9}))
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(twoSum([]int{2, 17, 7, 15}, 9))
}

/**
 *给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标
 */
func twoSum(nums []int, target int) []int {
	// 暴力枚举，时间复杂度：O(n²)（两层循环），空间复杂度：O(1)
	result := make([]int, 0)
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			result = append(result, i, j)
	//		}
	//	}
	//}

	// 使用哈希表，空间换时间：时间复杂度：O(n)（只需一次遍历），空间复杂度：O(n)（哈希表存储最多n个元素）
	m := make(map[int]int)
	for i, v := range nums {
		// 获取目标值减去当前值
		targetMinus := target - v
		// 判断map中是否存在目标值减去当前值
		if j, ok := m[targetMinus]; ok {
			// 存在则返回结果
			return []int{j, i}
		}
		// 不存在则将当前值添加到map中
		m[v] = i
	}
	return result
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 对intervals的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merge := [][]int{intervals[0]}
	// 逐个对比intervals其他区间集合，
	for i := 1; i < len(intervals); i++ {
		// merge最后一个区间
		last := merge[len(merge)-1]
		// 当前区间
		current := intervals[i]
		// 如果merge顶栈区间的结束区间小于当前区间的开始区间，则追加到 merge
		if last[1] < current[0] {
			merge = append(merge, intervals[i])
		} else {
			// 如果merge顶栈区间的结束区间大于当前区间的开始区间，则判断顶栈的结束区间是否小于当前区间的结束区间
			// 小于则更新顶栈区间的结束区间
			if last[1] < current[1] {
				last[1] = current[1]
			}
		}
	}

	return merge
}

/**
 * 删除数组重复元素,使用 O(1) 额外空间的条件下完成
 */
func rmDuplicates(ints []int) int {
	size := len(ints)
	i := 0
	newInts := []int{ints[0]}
	for j := 1; j < size-1; j++ {
		if ints[i] != ints[j] {
			i++
			ints[i] = ints[j]
			newInts = append(newInts, ints[i])
		}
	}
	fmt.Println(newInts)
	return i + 1
}

func plusOne(ints []int) []int {
	incr := 1
	size := len(ints)
	for i := size - 1; i >= 0; i-- {
		num := ints[i] + incr
		if num < 10 {
			ints[i] = num
			incr = 0
		} else {
			ints[i] = 0
		}
	}
	if incr == 1 {
		ints = append([]int{1}, ints...)
	}
	return ints
}

/**
 * 查找最长公共前缀
 */
func longestCommonPrefix(prefixes []string) string {
	first := prefixes[0]
	for i := 1; i < len(first); i++ {
		for j := 1; j < len(prefixes); j++ {
			if prefixes[j][i] != first[i] {
				return first[:i]
			}
		}
	}
	return first
}

/**
 * 验证括号
 */
func isValid(s string) bool {
	// 创建一个空栈
	stack := []rune{}
	// 创建一个括号对的 map集合
	maps := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	// 循环匹配字符串s
	for _, char := range s {
		switch char {
		// 判断是否左括号，是的话则入栈
		case '(', '[', '{':
			stack = append(stack, char)
		// 判断是否右括号，是的话，判断栈顶是否为空，为空则说明已经全部匹配，不为空则判断maps对应 key的值和栈顶的值是否匹配
		case ')', ']', '}':
			if len(stack) == 0 || maps[char] != stack[len(stack)-1] {
				return false
			}
			// 匹配则弹出栈顶
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

/**
 *	控制流程-只出现一次的数字
 */
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

/**
 * 控制流程-回文数：指正读和反读都相同的数字
 * 只反转数字的后半部分，然后与前半部分比较，这样可以避免完全反转整个数字带来的溢出风险，尤其是大数
 */
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
