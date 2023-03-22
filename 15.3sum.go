package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}

// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}

	var ret [][]int

	sort.Ints(nums)

	for i := 0; i < length-2; i++ {
		cur := nums[i]
		if cur > 0 {
			return ret
		}

		if i > 0 && cur == nums[i-1] {
			continue
		}

		l, r := i+1, length-1
		for l < r {
			left, right := nums[l], nums[r]
			x := cur + left + right
			if x == 0 {
				ret = append(ret, []int{cur, left, right})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if x > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return ret
}
