package main

import "fmt"

func main() {
	// arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	arr := []int{1, 100}
	fmt.Println(maxArea(arr))
}

// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0

	for left < right {
		area = Max(area, (right-left)*Min(height[left], height[right]))
		// 双指针, 矮的往内移动一步
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return area
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
