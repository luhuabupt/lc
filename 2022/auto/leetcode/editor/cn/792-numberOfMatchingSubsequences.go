package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func numMatchingSubseq(s string, words []string) (ans int) {
	n := len(s)

	dp := make([][26]int, n+1)
	for j := 0; j < 26; j++ {
		dp[n][j] = -1
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			dp[i][j] = dp[i+1][j]
		}
		dp[i][s[i]-'a'] = i
	}

loop:
	for _, v := range words {
		i := 0
		for _, x := range v {
			if dp[i][x-'a'] == -1 {
				continue loop
			}
			i = dp[i][x-'a'] + 1
		}
		ans++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
