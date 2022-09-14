package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func topoSort(a [][]int, all []int) []int {
	ans := []int{}

	in := map[int]int{} // 入度
	nx := map[int][]int{}
	for _, v := range a {
		in[v[1]]++

		if nx[v[0]] == nil {
			nx[v[0]] = []int{}
		}
		nx[v[0]] = append(nx[v[0]], v[1])
	}

	s := []int{} // 入度为0
	for _, v := range all {
		if in[v] == 0 {
			s = append(s, v)
		}
	}

	for len(s) > 0 {
		ns := []int{}
		for _, x := range s {
			ans = append(ans, x)
			for _, v := range nx[x] { // 拆边
				in[v]--
				if in[v] == 0 {
					ns = append(ns, v)
				}
			}
		}
		s = ns
	}

	if len(all) != len(ans) { // 有环
		return []int{}
	}

	return ans
}

func buildMatrix(k int, row [][]int, col [][]int) [][]int {
	all := make([]int, k)
	for i := 0; i < k; i++ {
		all[i] = i + 1
	}

	r := topoSort(row, all)
	c := topoSort(col, all)

	if len(r) == 0 || len(c) == 0 {
		return [][]int{}
	}

	pos := make([]int, k+1)
	for i, v := range c {
		pos[v] = i
	}

	ans := make([][]int, k)
	for i, v := range r {
		ans[i] = make([]int, k)
		ans[i][pos[v]] = v
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
