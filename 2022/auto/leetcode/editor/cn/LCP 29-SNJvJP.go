package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(orchestraLayout(5, i, j))
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func orchestraLayout(n int, x int, y int) int {
	minFour := func(a, b, c, d int) int {
		if a < b && a < c && a < d {
			return a
		}
		if b < c && b < d {
			return b
		}
		if c < d {
			return c
		}
		return d
	}

	l := minFour(x, y, n-1-x, n-1-y)
	s := l * (4*n - 4*l)
	t := n - 2*l - 1

	if x == l && y < n-1-l {
		//fmt.Println(1)
		s += y - l
	} else if y == n-1-l && x < n-1-l {
		//fmt.Println(2)
		s += t + x - l
	} else if x == n-1-l && y > 0 {
		//fmt.Println(3)
		s += 2*t + n - 1 - l - y
	} else {
		//fmt.Println(4)
		s += 3*t + n - 1 - l - x
	}

	return s%9 + 1
}

//leetcode submit region end(Prohibit modification and deletion)
