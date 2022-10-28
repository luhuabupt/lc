package main

import "fmt"

func main() {
	fmt.Println(arrangeBookshelf([]int{2, 2, 1, 2}, 2))
	fmt.Println(arrangeBookshelf([]int{2, 2, 4, 2}, 2))
	fmt.Println(arrangeBookshelf([]int{2, 1, 2, 2, 1, 3, 3, 1, 3, 3}, 2))
}

func arrangeBookshelf(a []int, limit int) []int {
	st := []int{}
	l := map[int]int{}
	c := map[int]int{}
	for _, v := range a {
		l[v]++
	}
	for _, v := range a {
		if c[v] == limit {
			l[v]--
			continue
		}
		for len(st) > 0 && v < st[len(st)-1] {
			x := st[len(st)-1]
			if l[x]+c[x] <= limit {
				break
			}

			c[x]--
			st = st[:len(st)-1]
		}

		c[v]++
		l[v]--
		st = append(st, v)
	}

	return st
}
