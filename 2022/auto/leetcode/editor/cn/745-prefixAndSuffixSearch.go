package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type WordFilter struct {
	m map[string]map[string][]int
}

func Constructor(words []string) WordFilter {
	m := map[string]map[string][]int{}

	for k, w := range words {
		n := len(w)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				pre := w[:i+1]
				suf := w[j:]
				if m[pre] == nil {
					m[pre] = map[string][]int{}
				}
				if m[pre][suf] == nil {
					m[pre][suf] = []int{}
				}

				m[pre][suf] = append(m[pre][suf], k)
			}
		}
	}

	return WordFilter{m}
}

func (t *WordFilter) F(pre string, suf string) int {
	if t.m[pre] == nil || t.m[pre][suf] == nil || len(t.m[pre][suf]) == 0 {
		return -1
	}
	return t.m[pre][suf][len(t.m[pre][suf])-1]
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
//leetcode submit region end(Prohibit modification and deletion)
