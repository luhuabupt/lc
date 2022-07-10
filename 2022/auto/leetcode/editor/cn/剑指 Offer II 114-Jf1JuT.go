package main

import "fmt"

func main() {
	fmt.Println(alienOrder([]string{"z", "x", "z"}))
	fmt.Println(alienOrder([]string{"z", "z"}))
	fmt.Println(alienOrder([]string{"za"}))
	fmt.Println(alienOrder([]string{"z", "za"}))
	fmt.Println(alienOrder([]string{"ac", "ab", "b"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func alienOrder(w []string) string {
	n := len(w)
	do := [26][26]bool{}
	in := [26]int{}
	has := [26]bool{}
	for _, v := range w {
		for _, x := range v {
			has[x-'a'] = true
		}
	}

loop:
	for i := 1; i < n; i++ {
		a, b := w[i-1], w[i]
		for j := 0; j < len(a) && j < len(b); j++ {
			if a[j] != b[j] {
				do[a[j]-'a'][b[j]-'a'] = true
				in[b[j]-'a']++
				continue loop
			}
		}
		if len(a) > len(b) {
			return ""
		}
	}

	ans := []byte{}
	no := false
	vis := [26]bool{}
	var dfs func(i int)
	dfs = func(i int) {
		ans = append(ans, byte('a'+i))
		vis[i] = true
		for j, x := range do[i] {
			if !x {
				continue
			}
			if vis[j] {
				no = true
				return
			}
			dfs(j)
		}
	}

	for i, inv := range in {
		if has[i] && inv == 0 {
			dfs(i)
		}
	}
	for i, v := range vis {
		if has[i] && !v {
			for _, x := range do[i] { // 有环
				if x {
					return ""
				}
			}
			ans = append(ans, byte('a'+i))
		}
	}
	if no {
		return ""
	}

	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
