package main

import "fmt"

func main() {
	//fmt.Println(numPermsDISequence("DID"))
	fmt.Println(numPermsDISequence("IDD"))
	//fmt.Println(numPermsDISequence("IIDIIDDIDDDDIIDDIDIDIDDDDIDDDIIIIDDIDDDDIDIIDDIDID"))
	//fmt.Println(numPermsDISequence("IIDII"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numPermsDISequence_(s string) int {
	n := len(s)
	dp := map[int]int{-1: 0}
	for i := 0; i <= n; i++ {
		dp[i] = i + 1
	}
	for _, v := range s {
		if v == 'D' {
			for i := 0; i < n; i++ {
				dp[i] = dp[n] - dp[i-1]
			}
		}
		if v == 'i' {
			for i := n; i > 0; i-- {
				dp[i] = dp[i-1]
			}
		}
		for i := 0; i <= n; i++ {
			dp[i] += dp[i-1]
		}
	}

	return dp[n]
}

func numPermsDISequence(s string) int {
	M := int(1E9) + 7
	vis := map[string]int{"D": 1, "I": 1, "": 1}


	cnt := 0
	var dfs func(s string) int
	dfs = func(s string) int {
		//fmt.Println(s)
		cnt++
		//fmt.Println(s)
		if vis[s] > 0 {
			return vis[s]
		}

		sum := 0
		for i := 0; i < len(s); i++ {
			x := s[i]

			if i == 0 && x == 'D' {
				sum += dfs(s[1:])
			}
			if i+1 < len(s) && x == 'I' && s[i+1] == 'D' {
				sum += c(n, i) * dfs(s[:i]) * dfs(s[i+2:])
				i++
			}
			if i == len(s)-1 && x == 'I' {
				sum += dfs(s[:len(s)-1])
			}

			sum %= M
		}
		vis[s] = sum % M

		return vis[s]
	}

	dfs(s)
	//fmt.Println(cnt)
	//fmt.Println(vis)

	return vis[s]
}

//leetcode submit region end(Prohibit modification and deletion)
