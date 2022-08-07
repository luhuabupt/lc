package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func possibleBipartition(n int, dis [][]int) bool {
	do := make([]int, n+1)
	for i := 1; i <= n; i++ {
		do[i] = -1
	}

	next := make([][]int, n+1)
	for _, v := range dis {
		next[v[0]] = append(next[v[0]], v[1])
		next[v[1]] = append(next[v[1]], v[0])
	}

	for i, id := range do {
		if id != -1 {
			continue
		}
		idx := 0
		st := []int{i}
		do[i] = idx

		for len(st) > 0 {
			idx ^= 1
			nt := []int{}
			for _, v := range st {
				for _, x := range next[v] {
					if do[x] == idx {
						continue
					} else if do[x] == -1 {
						do[x] = idx
						nt = append(nt, x)
					} else {
						return false
					}
				}
			}
			st = nt
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
