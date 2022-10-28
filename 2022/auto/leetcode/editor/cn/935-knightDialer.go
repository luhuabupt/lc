package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func knightDialer(n int) int {
	// 04, 06, 16, 18, 27, 29, 34, 38, 40, 43, 49, 60, 61, 67, 72, 76, 81, 83, 92, 94
	M := int(1e9) + 7
	d := [10][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{0, 3, 9},
		{},
		{0, 1, 7},
		{2, 6},
		{1, 3},
		{2, 4},
	}
	pre, cur := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, [10]int{}
	for i := 2; i <= n; i++ {
		cur = [10]int{}
		for i, v := range pre {
			for _, x := range d[i] {
				cur[x] += v
				cur[x] %= M
			}
		}

		pre = cur
	}

	ans := 0
	for _, v := range pre {
		ans += v
		ans %= M
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
