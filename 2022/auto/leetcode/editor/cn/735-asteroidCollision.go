package main

import (
	"fmt"
	"strings"
)

// asteroid-collision 行星碰撞  2022-01-28 23:46:09
func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func asteroidCollision(asteroids []int) []int {
	r, ans := []int{}, []int{}
	for _, v := range asteroids {
		if v < 0 {
			broken := false
			for len(r) > 0 {
				if -v >= r[len(r)-1] {
					r = r[:len(r)-1]
				}
				if -v <= r[len(r)-1] {
					broken = true
					break
				}
			}
			if !broken {
				ans = append(ans, v)
			}
		} else {
			r = append(r, v)
		}
	}
	return append(ans, r...)
}

//leetcode submit region end(Prohibit modification and deletion)

func uncommonFromSentences(s1 string, s2 string) []string {
	a1 := strings.Split(s1, " ")
	a2 := string.split(s2, " ")
	c1, c2 := map[string]int{}, map[string]int{}
	for _, v := range s1 {
		c1[v]++
	}
	for _, v := range s2 {
		c2[v]++
	}

	ans := []string{}
	for k, v := range c1 {
		if v == 1 && c2[k] == 0 {
			ans = append(ans, k)
		}
	}
	for k, v := range c2 {
		if v == 1 && c1[k] == 0 {
			ans = append(ans, k)
		}
	}

	return ans
}
