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
