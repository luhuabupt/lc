package main

func main() {
	// 2 8 10 18
}

//leetcode submit region begin(Prohibit modification and deletion)
func lenLongestFibSubseq(a []int) int {
	n := len(a)
	m := map[int]bool{}
	for _, v := range a {
		m[v] = true
	}

	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			tmp := 2
			for x, y := a[i], a[j]; m[x+y]; {
				x, y = y, x+y
				tmp++
			}
			if tmp > 2 && tmp > ans {
				ans = tmp
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
