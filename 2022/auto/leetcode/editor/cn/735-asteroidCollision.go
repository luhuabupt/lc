package main

import (
	"fmt"
)

// asteroid-collision 行星碰撞  2022-01-28 23:46:09
func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func asteroidCollision(a []int) []int {
	st := []int{}
	for _, v := range a {
		alive := true
		for v < 0 && len(st) > 0 && alive && st[len(st)-1] > 0 {
			alive = -v > st[len(st)-1]
			if -v >= st[len(st)-1] {
				st = st[:len(st)-1]
			}
		}
		if alive {
			st = append(st, v)
		}
	}
	return st
}

func asteroidCollision_(a []int) []int {
	r := []int{}
	ans := []int{}
	for _, v := range a {
		if v > 0 {
			r = append(r, v)
		} else {
			broken := false
			for len(r) > 0 && !broken {
				if -v <= r[len(r)-1] {
					broken = true
				}
				if -v >= r[len(r)-1] {
					r = r[:len(r)-1]
				}
			}
			if !broken {
				ans = append(ans, v)
			}
		}
	}

	return append(ans, r...)
}

//leetcode submit region end(Prohibit modification and deletion)
