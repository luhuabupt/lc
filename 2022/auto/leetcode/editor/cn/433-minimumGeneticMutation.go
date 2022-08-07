package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func minMutation(s string, e string, bank []string) int {
	flag := false
	for _, v := range bank {
		if e == v {
			flag = true
		}
	}
	if !flag {
		return -1
	}

	if s == e {
		return 0
	}

	n := len(bank)
	next := make([][]int, n)

	chk := func(a, b string) bool {
		cnt := 0
		for i := 0; i < 8; i++ {
			if a[i] != b[i] {
				cnt++
			}
		}
		return cnt == 1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if chk(bank[i], bank[j]) {
				next[i] = append(next[i], j)
				next[j] = append(next[j], i)
			}
		}
	}

	st := []int{}
	for i, v := range bank {
		if chk(s, v) {
			st = append(st, i)
		}
	}

	ans := 0
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = n + 1
	}

	for len(st) > 0 {
		ans++
		if ans > 9 {
			return -1
		}
		ns := []int{}
		for _, v := range st {
			if dp[v] < ans {
				continue
			}
			dp[v] = ans
			if bank[v] == e {
				return ans
			}

			for _, x := range next[v] {
				if dp[x] == n+1 {
					ns = append(ns, x)
				}
			}
		}
		st = ns
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
