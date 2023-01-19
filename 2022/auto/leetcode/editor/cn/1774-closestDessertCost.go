package main

import "fmt"

// closest-dessert-cost 最接近目标价格的甜点成本  2022-12-04 13:55:50
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
func closestCost(a []int, b []int, t int) int {
	m := int(3e4) + 2
	d := make([]bool, m)
	d[0] = true
	for _, v := range b {
		for x := m - 1; x >= 0; x-- {
			if d[x] {
				if x+v < m {
					d[x+v] = true
				}
				if x+v+v < m {
					d[x+2*v] = true
				}
			}
		}
	}

	e := make([]bool, m)
	for _, v := range a {
		for j, x := range d {
			if x {
				if v+j < m {
					e[v+j] = true
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		if t-i >= 0 && e[t-i] {
			return t - i
		}
		if t+i < m && e[t+i] {
			return t + i
		}
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
