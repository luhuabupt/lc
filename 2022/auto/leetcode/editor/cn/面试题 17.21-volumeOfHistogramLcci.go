package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(height)
	l, r := 0, make([]int, n)
	r[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		r[i] = max(r[i+1], height[i+1])
	}

	ans := 0
	for i, v := range height {
		ans += max(0, min(l, r[i])-v)
		l = max(v, l)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
