package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func minOperations(a []int, q []int) []int64 {
	sort.Ints(a)
	d := make([]int, len(a))
	t := 0
	for i, v := range a {
		t += v
		d[i] = t
	}

	ans := make([]int64, len(q))
	for i, v := range q {
		j := sort.SearchInts(a, v)
		x := j*v + t - (len(a)-j)*v
		if j > 0 {
			x -= d[j-1]
			x -= d[j-1]
		}
		ans[i] = int64(x)
	}

	return ans
}

func primeSubOperation(a []int) bool {
	// 所有质数
	prime := []int{0}
	d := make([]bool, 1001)
	for l := 2; l < 1001; l++ {
		for i := l + l; i < 1001; i += l {
			d[i] = true
		}
	}
	for i, v := range d {
		if i >= 2 && !v {
			prime = append(prime, i)
		}
	}

loop:
	for i, v := range a {
		for j := len(prime); j >= 0; j-- {
			if v-prime[j] > 0 && (i == 0 || v-prime[j] > a[i-1]) {
				a[i] = v - prime[j]
				continue loop
			}
		}
		return false
	}

	return true
}
