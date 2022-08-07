package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	lt := map[int]map[int]bool{} // 亮
	r := map[int]int{}           // 行
	c := map[int]int{}           // 列
	l := map[int]int{}           // 对角线
	rl := map[int]int{}          // 反对角线
	dir := [][2]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for _, v := range lamps {
		if lt[v[0]] == nil {
			lt[v[0]] = map[int]bool{}
		}
		if lt[v[0]][v[1]] {
			continue
		}
		lt[v[0]][v[1]] = true
		r[v[0]]++
		c[v[1]]++
		l[n-1+v[1]-v[0]]++
		rl[v[0]+v[1]]++
	}

	var ans []int
	for _, v := range queries {
		i, j := v[0], v[1]
		if r[i] > 0 || c[j] > 0 || l[n-1+j-i] > 0 || rl[i+j] > 0 {
			ans = append(ans, 1)
		} else {
			ans = append(ans, 0)
		}

		for _, d := range dir {
			a, b := i+d[0], j+d[1]
			if a >= 0 && a < n && b >= 0 && b < n {
				if lt[a] != nil && lt[a][b] { // 亮
					r[a]--
					c[b]--
					l[n-1+b-a]--
					rl[a+b]--
					lt[a][b] = false
				}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
