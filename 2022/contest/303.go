package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countExcellentPairs([]int{1<<31 - 1, 1, 3}, 30))
}

func countExcellentPairs(a []int, k int) int64 {
	m := map[int]bool{}
	for _, v := range a {
		m[v] = true
	}
	ans := 0
	c := [32]int{}
	for v, _ := range m {
		c[bits.OnesCount(uint(v))]++
	}
	n := len(m)
	for v, _ := range m {
		x := bits.OnesCount(uint(v))
		if x >= k {
			ans += n
			continue
		}
		for i := k - x; i < 32; i++ {
			ans += c[i]
		}
	}
	return int64(ans)
}
