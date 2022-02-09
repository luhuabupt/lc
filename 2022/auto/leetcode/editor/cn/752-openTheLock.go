package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func openLock(deadends []string, target string) int {
	dp := map[string]int{}

	m := map[string]bool{}
	for _, v := range deadends {
		if v == "0000" {
			return -1
		}
		m[v] = true
	}

	st, idx := []string{"0000"}, 0
	for idx < len(st) {
		cur := st[idx]
		if cur == target {
			return dp[cur]
		}

		for i, v := range cur {
			x := []byte(cur)

			x[i] = byte((v-'0'+1)%10 + '0')
			if !m[string(x)] && dp[string(x)] == 0 {
				st = append(st, string(x))
				dp[string(x)] = dp[cur] + 1
			}

			x[i] = byte((v-'0'+9)%10 + '0')
			if !m[string(x)] && dp[string(x)] == 0 {
				st = append(st, string(x))
				dp[string(x)] = dp[cur] + 1
			}
		}
		idx++
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
