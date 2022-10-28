package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countSpecialNumbers(20))
	fmt.Println(countSpecialNumbers(5))
	fmt.Println(countSpecialNumbers(133))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	ans := 0
	dp := map[int]int{-1: 0}

	flag := true
	c := [10]int{}
	for i, v := range s {
		//fmt.Println(i, string(v), c)
		if flag {
			for j := 0; j < int(v-'0'); j++ {
				//fmt.Println(j, v-'0', c[v-'0'])
				if i == 0 && j == 0 {
					continue
				}
				if c[j] == 0 {
					dp[i]++
				}
			}
		}

		dp[i] += dp[i-1] * (10 - i)

		c[v-'0']++
		if c[v-'0'] > 1 {
			flag = false
		}
	}

	//fmt.Println(dp)

	for i := 0; i < len(s)-1; i++ {
		t := 9
		for j := 9; j >= 10-i; j-- {
			t *= j
		}
		ans += t
	}
	if flag {
		ans++
	}

	return ans + dp[len(s)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
