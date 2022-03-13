package main

import "sort"

func minEatingSpeed(p []int, h int) int {
	return sort.Search(1E9, func(i int) bool {
		x := 0
		for _, v := range p {
			x += v / i
			if v % i > 0 {
				x++
			}
		}

		return x <= h
	})
}
