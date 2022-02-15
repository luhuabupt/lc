package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(wordBreak("catskicatcats", []string{"cats", "cat", "dog", "ski"}))
	//fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, w []string) bool {
	sort.Slice(w, func(i, j int) bool {
		return len(w[i]) < len(w[j])
	})
	vis := make([]bool, len(s)+1)

	var dfs func(v string) bool
	dfs = func(v string) bool {
		for _, x := range w {
			if len(x) > len(v) {
				break
			}
			if x == v || v[:len(x)] == x && !vis[len(v)-len(x)] && dfs(v[len(x):]) {
				return true
			}
		}
		vis[len(v)] = true
		return false
	}

	return dfs(s)
}

//leetcode submit region end(Prohibit modification and deletion)
