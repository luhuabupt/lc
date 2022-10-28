package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func minIncrementForUnique(a []int) int {
	ans := 0
	m := map[int]int{}
	for _, v := range a {
		m[v]++
	}
	for i := 0; i < 2e5+2; i++ {
		if m[i] > 1 {
			ans += m[i] - 1
			m[i+1] += m[i] - 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
