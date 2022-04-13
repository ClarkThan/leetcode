package main

import "fmt"

func main() {
	// arr := []int{0, 1, 1, 1, 1, 1, 2, 2, 3, 3, 4}
	arr := []int{1}
	fmt.Println(removeDuplicates(arr))
}

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	idx := 0
	for i, n := range nums {
		if n != nums[idx] {
			idx++
			// 保证相邻的两个不同元素的情况下不需要做修改
			if i > idx {
				nums[idx] = n
			}
		}
	}

	return idx + 1
}

func removeDuplicatesSlowFast(nums []int) int {
	slow, fast, length := 0, 1, len(nums)
	if length <= 1 {
		return length
	}

	for {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
			fast++
		}

		if fast >= length {
			break
		}
	}

	return slow + 1
}
