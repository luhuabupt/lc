package main

import "fmt"

func main() {
	//fmt.Println(racecar(1))
	//fmt.Println(racecar(2))
	//fmt.Println(racecar(3))
	fmt.Println(racecar(4))
	fmt.Println(racecar(5))
	fmt.Println(racecar(10000))
}

//leetcode submit region begin(Prohibit modification and deletion)
func racecar(n int) int {
	dp := make(map[int]map[int]int)
	dp[0] = map[int]int{}
	dp[0][1] = 0

	st, nx, v := [][2]int{{0, 1}}, 0, 1
	for len(st) > 0 {
		cur := st[0]
		ci, cv := cur[0], cur[1]
		st = st[1:]

		// A
		nx, v = ci+cv, cv*2
		if dp[nx] == nil {
			dp[nx] = map[int]int{}
		}
		if dp[nx][v] == 0 {
			dp[nx][v] = dp[ci][cv] + 1
			if nx == n {
				//fmt.Println(dp)
				return dp[nx][v]
			}
			st = append(st, [2]int{nx, v})
		}

		// R
		if cv > 0 {
			v = -1
		} else {
			v = 1
		}
		if _, ok := dp[ci][v]; !ok {
			dp[ci][v] = dp[ci][cv] + 1
			st = append(st, [2]int{ci, v})
		}
	}

	return 1e5
}

//leetcode submit region end(Prohibit modification and deletion)
