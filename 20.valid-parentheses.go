package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("([]{})"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid(")("))
	fmt.Println(isValid("())"))
}

// https://leetcode.cn/problems/valid-parentheses/?favorite=2cktkvj
func isValid(s string) bool {
	m := map[byte]byte{'(': ')', '{': '}', '[': ']'}

	stack := make([]byte, 0, len(s))
	for _, c := range []byte(s) {
		if _, exists := m[c]; exists {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return false
		} else {
			// if m[stack.pop()] != c
			x := stack[len(stack)-1]
			if m[x] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
