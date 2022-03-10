package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func robotSim(c []int, o [][]int) (ans int) {
	do := map[int]map[int]bool{}

	for _, a := range o {
		if do[a[0]] == nil {
			do[a[0]] = map[int]bool{}
		}
		do[a[0]][a[1]] = true
	}

	dir := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, idx := 0, 0, 0
	for _, v := range c {
		if v == -1 {
			idx = (idx + 1) % 4
		} else if v == -2 {
			idx = (idx + 3) % 4
		} else {
			d := dir[idx]
			for i := 0; i < v; i++ {
				if do[x+d[0]][y+d[1]] {
					break
				}
				x += d[0]
				y += d[1]
				if x*x+y*y > ans {
					ans = x*x + y*y
				}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
