package main

import "strconv"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	ans := []string{}

	var dfs func(cur []byte, l, i, use int) // [l, i]
	dfs = func(cur []byte, l, i, use int) {
		if use == 4 && i == len(s)-1 {
			ans = append(ans, string(cur))
			return
		}

		if i-l == 2 {
			x, _ := strconv.Atoi(s[l : i+1])
			if x > 255 {
				return
			}
		}
		if i != l && s[l] == '0' {
			return
		}

		dfs(cur, l, i+1, use)

		cur = append(cur, s[l:i+1]...)
		cur = append(cur, '.')
		dfs(cur, i+1, i+1, use+1)
		cur = cur[:l]
	}

	dfs([]byte{}, 0, 0, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
