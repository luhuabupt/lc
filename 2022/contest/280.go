package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(maximumANDSum([]int{14, 7, 9, 8, 2, 4, 11, 1, 9}, 8))
	fmt.Println(maximumANDSum([]int{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println(maximumANDSum([]int{1, 3, 10, 4, 7, 1}, 9))
	fmt.Println(maximumANDSum([]int{8, 1, 6, 14, 4, 7, 10, 11, 1, 4, 10, 10, 6, 8, 7, 10, 2, 10}, 9))
	fmt.Println(maximumANDSum([]int{8, 12, 12, 8, 9, 7, 13, 5, 15, 5, 6, 7, 2, 7}, 7)) // 51
}

func minimumRemoval(b []int) int64 {
	n := len(b)
	sort.Ints(b)
	s := make([]int, n)
	x := 0
	for i, v := range b {
		x += v
		s[i] = x
	}

	ans := int64(math.MaxInt64)

	for i := 0; i < n; i++ {
		t := int64(0)
		t += int64(s[n-1])
		t -= int64(b[i]) * int64(n-i)
		if t < ans {
			ans = t
		}
	}

	return ans
}

func maxNumberOfBalloons(text string) int {
	cnt := [26]int{}
	for _, v := range text {
		cnt[v-'a']++
	}
	cnt['l'-'a'] /= 2
	cnt['o'-'a'] /= 2
	ans := len(text)
	for _, v := range "balloon" {
		if cnt[v-'a'] < ans {
			ans = cnt[v-'a']
		}
	}
	return ans
}

func maximumANDSum(a []int, b int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)
	mask := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		mask[i] = -1
	}

	for i := 0; i < b; i++ {
		pre := append([]int{}, mask...)
		for k, v := range pre {
			if v == -1 {
				continue
			}
			oc := bits.OnesCount(uint(k))
			if oc < n-(b-i)*2 {
				continue
			} // 剪枝

			for j := 0; j < n; j++ { // 选1个
				if k>>j&1 == 0 {
					nk := k + (1 << j)
					mask[nk] = max(mask[nk], v+(a[j]&(i+1)))
					for z := 0; z < n; z++ { // 选第2个
						if nk>>z&1 == 0 {
							mask[nk+(1<<z)] = max(mask[nk+(1<<z)], v+(a[j]&(i+1))+(a[z]&(i+1)))
						}
					}
				}
			}
		}
	}
	return mask[1<<n-1]
}

func maximumANDSum_(a []int, b int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)
	mask := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		mask[i] = -1
	}

	for i := 0; i < b+b; i++ {
		for k, v := range mask {
			if v == -1 {
				continue
			}
			oc := bits.OnesCount(uint(k))
			if oc < n-(2*b-i) {
				continue
			} // 剪枝

			for j := 0; j < n; j++ {
				if k>>j&1 == 0 {
					mask[k|1<<j] = max(mask[k|1<<j], v+(a[j]&(i/2+1)))
				}
			}
		}
	}

	return mask[1<<n-1]
}
