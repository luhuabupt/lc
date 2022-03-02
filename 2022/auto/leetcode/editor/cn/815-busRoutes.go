package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	dp := make([]int, 1e6+1)
	for i, _ := range dp {
		dp[i] = -1
	}
	dp[source] = 0

	way := make([][]int, 1e6+1) // 点上的线路
	for k, v := range routes {
		for _, x := range v {
			way[x] = append(way[x], k)
		}
	}
	visWay := make([]bool, 500)

	st := []int{source}
	for len(st) > 0 {
		cur := st[0]
		st = st[1:]

		// 上哪条线
		for _, w := range way[cur] {
			if !visWay[w] {
				visWay[w] = true
			} else {
				continue
			}

			for _, x := range routes[w] { // 线路上所有站
				if x == target {
					return dp[cur] + 1
				}
				if dp[x] != -1 {
					continue
				}
				if len(way[x]) > 1 { // 换乘点
					st = append(st, x)
				}
				dp[x] = dp[cur] + 1
			}
		}
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
