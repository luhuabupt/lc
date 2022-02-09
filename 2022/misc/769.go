package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}

func maxChunksToSorted(arr []int) (ans int) {
	max := 0
	for i, v := range arr {
		if v > max {
			max = v
		}
		if i == max {
			ans++
		}
	}

	return
}

func maxChunksToSorted_(arr []int) (ans int) {
	n := len(arr)
	sp := make([][]bool, n)
	for i := 0; i < n; i++ {
		sp[i] = make([]bool, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i < n-l; i++ {
			mi, ma := 100, -1
			for x := i; x <= i+l; x++ {
				if arr[x] < mi {
					mi = arr[x]
				}
				if arr[x] > ma {
					ma = arr[x]
				}
			}
			if mi == i && ma == i+l {
				sp[i][i+l] = true
			}
		}
	}

loop:
	for i := 0; i < (1 << (n - 1)); i++ {
		pre := 0
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				if !sp[pre][j] {
					continue loop
				}
				pre = j + 1
			}
		}
		x := bits.OnesCount(uint(i))
		if x > ans {
			ans = x
		}
	}

	return ans + 1
}
