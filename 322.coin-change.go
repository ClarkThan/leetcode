package main

import "fmt"

func main() {
	// coins := []int{2}
	// amount := 3
	coins := []int{186, 419, 83, 408}
	amount := 6249
	fmt.Println(coinChangeRec(coins, amount))
}

// https://leetcode-cn.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// for i := 1; i <= amount; i++ {
	// 	for _, c := range coins {
	// 		if i >= c {
	// 			dp[i] = min(dp[i], dp[i-c]+1)
	// 		}
	// 	}
	// }
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}

	ret := dp[amount]
	if ret > amount {
		return -1
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChangeRec(coins []int, amount int) int {
	m := map[int]int{-1: amount + 1, 0: 0}
	ret := helper(m, coins, amount)
	if ret > amount {
		return -1
	}
	return ret
}

func helper(m map[int]int, coins []int, amount int) int {
	v, exist := m[amount]
	if exist {
		return v
	}

	// min(dp[amount - coins[0]], dp[amount - coins[1]], ...)
	minCnt := m[-1]
	for _, c := range coins {
		extra := amount - c
		if extra >= 0 {
			tmp := helper(m, coins, extra)
			if tmp < minCnt {
				minCnt = tmp
			}
		}
	}

	// if minCnt > amount {
	// 	m[amount] = minCnt
	// 	return minCnt
	// } else {
	// 	ret := minCnt + 1
	// 	m[amount] = ret
	// 	return ret
	// }

	ret := minCnt + 1
	m[amount] = ret
	return ret
}
