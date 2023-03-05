package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1}
	// nums = []int{4, 1, 2, 1, 2}
	// nums = []int{1}
	fmt.Println(singleNumber(nums))
}

// https://leetcode.cn/problems/single-number/?favorite=2cktkvj
func singleNumber(nums []int) int {
	ret := 0
	for _, n := range nums {
		ret ^= n
	}

	return ret
}
