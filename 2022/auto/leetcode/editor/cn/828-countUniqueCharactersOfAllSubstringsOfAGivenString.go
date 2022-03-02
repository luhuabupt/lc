package main

import "fmt"

func main() {
	fmt.Println(uniqueLetterString("ABC"))      // 10
	fmt.Println(uniqueLetterString("ABA"))      // 8
	fmt.Println(uniqueLetterString("LEETCODE")) // 92

	//fmt.Println(uniqueLetterString("CAABA")) // 19
}

//leetcode submit region begin(Prohibit modification and deletion)
func uniqueLetterString(s string) int {
	n := len(s)
	M := int(1e9) + 7
	dp := make([]int, n)
	pre := make([][2]int, n)

	a := [26]int{}
	for i, _ := range a {
		a[i] = -1
	}
	for i, _ := range pre {
		pre[i] = [2]int{-1, -1}
	}
	for i, v := range s {
		pre[i][1] = a[v-'A']
		if a[v-'A'] != -1 {
			pre[i][0] = pre[a[v-'A']][1]
		}
		a[v-'A'] = i
	}
	//fmt.Println(pre)

	ans := 0
	for i, _ := range s {
		if i == 0 {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + i + 1
			dp[i] += M - 2*pre[i][1] + pre[i][0] - 1
			dp[i] %= M
		}

		ans += dp[i]
		ans %= M
	}

	//fmt.Println(dp)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
