package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func minMoves(a []int, k int) int {
	// k
	// 奇 - a[i] = 1 左k/2 右k/2
	// 偶 - a[i] = 1 左k/2-1 右k/2
	//    - a[i] = 0 左k/2 右k/2

	t, t1 := 0, 0
	for i := 1; i < k/2; i++ {
		t1 += i
	}
	t = t1 + k/2

	n := len(a)
	l := make([]int, n)
	l1 := make([]int, n)
	r := make([]int, n)
	for i, _ := range l {
		l[i], l1[i], r[i] = -1, -1, -1
	}

	c, sc, il, ir := 0, 0, 0, 0
	for c < k/2 {
		ir++
		c += a[ir]
	}
	for ir < n {

	}
}

//leetcode submit region end(Prohibit modification and deletion)
