package main

// letter-case-permutation 字母大小写全排列  2022-02-15 20:37:19
func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func letterCasePermutation(s string) (ans []string) {
	n := len(s)

	var dfs func(pre string, i int)
	dfs = func(pre string, i int) {
		if i == n {
			ans = append(ans, pre)
		}
		dfs(pre+string(s[i]), i+1)
		if s[i] >= 'a' && s[i] <= 'z' {
			dfs(pre+string(s[i]+'A'-'a'), i+1)
		}
	}

	dfs("", 0)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
