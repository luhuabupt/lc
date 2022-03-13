package main

import "sort"

func nthMagicalNumber(n int, a int, b int) int {
	if a < b {
		return nthMagicalNumber(n, b, a)
	}

	gcd := func(i, j int) int {
		for i > 0 {
			i, j = i%j, i
		}
		return j
	}
	lcm := a * b / gcd(a, b)

	ai := sort.Search(a*1E9+1, func(i int) bool {
		return i/a+i/b-i/lcm >= n
	})

	return ai * a
}
