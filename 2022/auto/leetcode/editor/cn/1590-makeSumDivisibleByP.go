package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func minSubarray(a []int, p int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	s %= p
	if s == 0 {
		return 0
	}

	m := map[int]int{0: -1}
	ans := 1 << 30
	t := 0
	for i, v := range a {
		t += v
		t %= p

		x := (p + t - s) % p
		if _, ok := m[x]; ok {
			if i-m[x] < ans {
				ans = i - m[x]
			}
		}
		m[t] = i
	}

	if ans == 1<<30 || ans == len(a) {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
