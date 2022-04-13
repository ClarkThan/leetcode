package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

// https://leetcode-cn.com/problems/two-sum
func twoSum(nums []int, target int) []int {
	// 当前元素差值 -> 当前元素索引
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		diff, exist := m[n]
		if exist {
			return []int{diff, i}
		}
		m[target-n] = i
	}
	return nil
}
