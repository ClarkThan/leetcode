package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abzzzzba"))
}

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// left 不动，i 向后移动
	// 当 i 遇到重复字符，left 应该放在上一个重复字符的位置的后一位，同时记录最长的长度
	// 怎样判断是否遇到重复字符，且怎么知道上一个重复字符的位置？--用哈希字典的key来判断是否重复，用value来记录该字符的下一个不重复的位置。

	windowMap := make(map[byte]int)
	var max, left int
	for i, c := range []byte(s) {
		idx, exist := windowMap[c]
		if exist {
			left = Max(idx+1, left)
		}
		max = Max(i-left+1, max)
		windowMap[c] = i
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
