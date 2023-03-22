package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
}

var ret []string

var d2c = map[byte][]byte{
	'2': []byte("abc"),
	'3': []byte("def"),
	'4': []byte("ghi"),
	'5': []byte("jkl"),
	'6': []byte("mno"),
	'7': []byte("pqrs"),
	'8': []byte("tuv"),
	'9': []byte("wxyz"),
}

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	ret = nil // 清空全局变量
	backtrack(digits, 0, "")
	return ret
}

func backtrack(digits string, idx int, combination string) {
	if idx == len(digits) {
		ret = append(ret, combination)
	} else {
		for _, c := range d2c[digits[idx]] {
			tmp := string(append([]byte(combination), c))
			backtrack(digits, idx+1, tmp)
		}
	}
}
