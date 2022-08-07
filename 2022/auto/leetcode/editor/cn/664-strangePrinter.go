package main

import "fmt"

// strange-printer 奇怪的打印机  2022-01-09 21:55:01
func main() {
	fmt.Println(strangePrinter("aba"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func strangePrinter_(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			if l == 1 {
				dp[i][i] = 1
				continue
			}
			max := 0
			for j := i; j+1 <= i+l-1; j++ {
				x := dp[i][j] + dp[j+1][i+l-1]
				if x > max {
					max = x
				}
			}
			dp[i][i+l-1] = max
		}
	}
	return dp[0][n-1]
}

func strangePrinter(s string) int {
	arr := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] != arr[len(arr)-1] {
			arr = append(arr, s[i])
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := 0; i < n-1; i++ {
		dp[i][i+1] = 1
		if arr[i] != arr[i+1] {
			dp[i][i+1]++
		}
	}
	for l := 3; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			dp[i][i+l-1] = min(dp[i][i+l-2], dp[i+1][i+l-1])
			if arr[i] != arr[i+l-1] {
				dp[i][i+l-1]++
			}
		}
	}
	return dp[0][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
