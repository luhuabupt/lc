package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func pyramidTransition(bottom string, allowed []string) bool {
	m := make([][]int, 64)
	vis := map[int]bool{}
	for _, v := range allowed {
		x := int(v[0]-'A'+1) + 7*int(v[1]-'A'+1)
		m[x] = append(m[x], int(v[2]-'A'+1))
	}

	var dfs func(s int) bool
	dfs = func(s int) bool {

	}

	s, p := 0, 1
	for _, v := range bottom {
		s += int(v-'A'+1) * p
		p *= 8
	}

	return dfs(s)
}

//leetcode submit region end(Prohibit modification and deletion)
