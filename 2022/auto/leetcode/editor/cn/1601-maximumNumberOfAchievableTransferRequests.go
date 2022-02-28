package main

import "fmt"

func main() {
	fmt.Println(maximumRequests(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumRequests(m int, r [][]int) int {
	n := len(r)
	ans := 0

loop:
	for i := 0; i < 1<<n; i++ {
		vis := make([]int, m)
		cnt := 0

		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				vis[r[j][0]]--
				vis[r[j][1]]++
				cnt++
			}
		}
		for _, x := range vis {
			if x != 0 {
				continue loop
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
