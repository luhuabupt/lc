package main

import "fmt"

func main() {
	fmt.Println(reachingPoints(1, 4, 999999997, 4))
	fmt.Println(reachingPoints(18, 50, 111, 82))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	max := func(a, b, c int) int {
		if a > b && a > c {
			return a
		}
		if b > c {
			return b
		}
		return c
	}

	for {
		if ty == sy && sx == tx {
			return true
		}
		if ty < sy || tx < sx {
			return false
		}
		if ty > tx {
			x := max(tx, sx, sy)
			if ty-10*tx > x {
				ty = ty%tx + ((x-1)/tx+1)*tx
			} else {
				ty -= tx
			}
		} else {
			x := max(ty, sx, sy)
			if tx-10*ty > x {
				tx = tx%ty + ((x-1)/ty+1)*ty
				//fmt.Println(tx)
			} else {
				tx -= ty
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
