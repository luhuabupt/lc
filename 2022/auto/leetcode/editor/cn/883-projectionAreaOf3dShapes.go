package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func projectionArea(g [][]int) (ans int) {
	for _, a := range g {
		ma := 0
		for _, v := range a {
			if v > 0 {
				ans++
			}
			if v > ma {
				ma = v
			}
		}
		ans += ma
	}
	for j := 0; j < len(g); j++ {
		ma := 0
		for i := 0; i < len(g); i++ {
			if g[i][j] > ma {
				ma = g[i][j]
			}
		}
		ans += ma
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
