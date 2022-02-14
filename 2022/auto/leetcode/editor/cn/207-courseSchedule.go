package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(n int, a [][]int) bool {
	in, next := make([]int, n), make([][]int, n)
	for _, v := range a {
		in[v[1]]++
		next[v[0]] = append(next[v[0]], v[1])
	}
	s := []int{}
	for i, v := range in {
		if v == 0 {
			s = append(s, i)
		}
	}

	cnt := 0
	for len(s) > 0 {
		cnt += len(s)
		ns := []int{}
		for _, v := range s {
			if len(next[v]) == 0 {
				continue
			}
			for _, x := range next[v] {
				in[x]--
				if in[x] == 0 {
					ns = append(ns, x)
				}
			}
		}
		s = ns
	}

	return cnt == n
}

//leetcode submit region end(Prohibit modification and deletion)
