package main

import (
	"fmt"
)

func main() {
	s := "abcdedcba"
	s = "aaba"
	s = "bb"
	// s = "aaabbbbb"
	fmt.Println(longestPalindrome(s))
}

// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	sLen := len(s)
	dp := make([][]bool, sLen) // 二位数组, dp[i][j] 表示从位置 i到 j 是否满足回文
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}

	maxLen := 1
	si, ei := 0, 0
	for r := 1; r < sLen; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r+1-l > maxLen {
					maxLen = r + 1 - l
					si = l
					ei = r
				}
			}
		}
	}

	return s[si : ei+1]
}

func longestPalindromeV0(s string) string {
	sLen := len(s)
	if sLen < 2 {
		return s
	}

	var maxLen int
	var si, ei int

	for i := range s {
		length := 1
		left, right := i, i
		c := s[i]
		for left > 0 && s[left-1] == c {
			left--
			length++
		}
		for right < sLen-1 && s[right+1] == c {
			right++
			length++
		}
		for left > 0 && right < sLen-1 && s[left-1] == s[right+1] {
			left--
			right++
			length += 2
		}

		if length > maxLen {
			maxLen = length
			si, ei = left, right
		}
	}

	return s[si : ei+1]
}

// fmt.Println("max len", maxLen)
