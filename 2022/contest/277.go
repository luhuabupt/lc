package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(maximumGood([][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}}))
}

func maximumGood(statements [][]int) int {
	n := len(statements)
	ans := 0
	chk := func(i int) bool {
		for x := 0; x < n; x++ {
			if i&(1<<x) > 0 { // 好人
				for ss, v := range statements[x] {
					if v == 1 && (i&(1<<ss) == 0) { // 好
						return false
					}
					if v == 0 && (i&(1<<ss) > 0) { // 坏
						return false
					}
				}
			}
		}
		return true
	}

	for i := (1 << n) - 1; i >= 0; i-- {
		v := 0
		for x := 0; x < n; x++ {
			if (1<<x)&i > 0 {
				v += 1 << x
			}
		}
		if chk(v) && bits.OnesCount(uint(v)) > ans {
			ans = bits.OnesCount(uint(v))
		}
	}
	return ans
}
