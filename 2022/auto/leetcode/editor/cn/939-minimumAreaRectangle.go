package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func minAreaRect(points [][]int) int {
	n := len(points)
	m := map[int]map[int]bool{}
	for _, v := range points {
		if m[v[0]] == nil {
			m[v[0]] = map[int]bool{}
		}
		m[v[0]][v[1]] = true
	}

	mi := 1 << 60
	for i, v := range points {
		for j := i + 1; j < n; j++ {
			a, b, c, d := v[0], v[1], points[j][0], points[j][1]
			if a == c || b == d {
				continue
			}

			if m[a][d] && m[c][b] {
				s := (a - c) * (b - d)
				if s < 0 {
					s = -s
				}
				if s < mi {
					mi = s
				}
			}
		}
	}

	if mi == (1 << 60) {
		return 0
	}
	return mi
}

//leetcode submit region end(Prohibit modification and deletion)
