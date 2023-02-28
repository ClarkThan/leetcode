package main

import (
	"fmt"
)

// https://leetcode.cn/problems/product-of-array-except-self/
// 先计算所有的乘积, 再迭代除以对应元素的思路是错误的, 因为会有 0 元素的情况
// 第一次遍历: 先计算每个元素之前的所有元素的乘积
// 第二次遍历: 从后往前依次乘以后面的元素(上一次遍历没有乘的元素)
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}

	tail := 1
	for i := len(ret) - 2; i >= 0; i-- {
		tail *= nums[i+1]
		ret[i] = ret[i] * tail
	}

	return ret
}

func main() {
	fmt.Println(productExceptSelf([]int{1}))
	fmt.Println(productExceptSelf([]int{1, 2}))
	fmt.Println(productExceptSelf([]int{1, 2, 3}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
