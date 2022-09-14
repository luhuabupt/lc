package main

import "fmt"

func main() {
	fmt.Println(maximumSegmentSum([]int{1, 2, 5, 6, 1}, []int{0, 3, 2, 4, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumSegmentSum(a []int, rm []int) []int64 {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	n := len(a)
	ans := make([]int64, n)

	fa := make([]int, n)
	sum := make([]int, n)
	for i, _ := range fa {
		fa[i] = -1
	}

	var find func(i int) int
	find = func(i int) int {
		if fa[i] == i {
			return i
		}
		return find(fa[i])
	}

	union := func(i, j int) {
		n, o := find(i), find(j)
		fa[o] = n
		sum[n] += sum[o]
	}

	ma := 0
	for i := n - 1; i >= 0; i-- {
		ri := rm[i]
		fa[ri] = ri
		sum[ri] = a[ri]

		if ri > 0 && fa[ri-1] != -1 {
			union(ri, ri-1)
		}
		if ri < n-1 && fa[ri+1] != -1 {
			union(ri, ri+1)
		}

		ans[i] = int64(ma)
		ma = max(ma, sum[find(ri)])
	}

	fmt.Println(fa, sum)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
