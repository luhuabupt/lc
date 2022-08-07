package main

import (
	"fmt"
)

func main() {
	// 25
	fmt.Println(minimumTime("011001111111101001010000001010011"))
	fmt.Println(minimumTime("011001111111101"))
	fmt.Println(minimumTime("1010011"))
	fmt.Println(minimumTime("0100101000000"))
	fmt.Println(minimumTime("1100101"))
	fmt.Println(minimumTime("110101"))
	fmt.Println(minimumTime("0010"))
	fmt.Println(minimumTime("0"))
	fmt.Println(minimumTime("11"))
	fmt.Println(minimumTime("111"))
	fmt.Println(minimumTime("1"))
	fmt.Println(minimumTime("0"))
}

func minimumTime(s string) int {
	n := len(s)
	r := make([]int, n+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := n - 1; i >= 0; i-- {
		r[i] = r[i+1]
		if s[i] == '1' {
			r[i] = min(r[i]+2, n-i)
		}
	}

	ans, pre := n, 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			pre = min(pre+2, i+1)
		}
		ans = min(ans, pre+r[i+1])
	}

	return ans
}
