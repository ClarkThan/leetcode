package main

import "fmt"

func main() {
	fmt.Println(isMatch("abcd", "ab.dc*"))
	fmt.Println(isMatch("aa", "a*"))
}

// https://leetcode.cn/problems/regular-expression-matching/?favorite=2cktkvj
func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)

	match := func(i, j int) bool {
		if i == 0 {
			return false
		}

		if p[j-1] == '.' {
			return true
		}

		return s[i-1] == p[j-1]
	}

	dp := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true

	for i := 0; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] == '*' {
				if match(i, j-1) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				dp[i][j] = match(i, j) && dp[i-1][j-1]
			}
		}
	}

	return dp[sLen][pLen]
}
