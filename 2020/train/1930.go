package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countPalindromicSubsequence("aabca"))
}

// 1930. 长度为 3 的不同回文子序列
// @link https://leetcode-cn.com/problems/unique-length-3-palindromic-subsequences/
func countPalindromicSubsequence(s string) int {
	left, right := make([]int, len(s)), make([]int, len(s))
	left[0], right[len(s)-1] = 1<<(s[0]-97), 1<<(s[len(s)-1]-97)
	for i := 1; i < len(s); i++ {
		left[i] = left[i-1] | 1<<(s[i]-97)
	}
	for i := len(s) - 2; i >= 0; i-- {
		right[i] = right[i+1] | 1<<(s[i]-97)
	}
	m, ans := map[string]bool{}, 0
	for i := 1; i < len(s)-1; i++ {
		tmp := left[i-1] & right[i+1]
		for j := 0; j < 26; j++ {
			if (tmp & (1 << j)) > 0 {
				key := strconv.Itoa(j+97) + string(s[i])
				if !m[key] {
					m[key] = true
					ans++
				}
			}
		}
	}
	return ans
}
