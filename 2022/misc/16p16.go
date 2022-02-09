package main

import "fmt"

func main() {
	fmt.Println(subSort([]int{1, 2, 4, 4, 10, 11, 4, 12}))
}

func subSort(array []int) []int {
	n := len(array)
	if n == 0 {
		return []int{-1, -1}
	}
	min, max, l, r := 1<<31, -1<<31, 0, n-1
	for i := n - 1; i >= 0; i-- {
		v := array[i]
		if v > min {
			l = i
		} else if v == min {
			if array[i+1] == v {

			} else {
				l = i + 1
			}
		}
		if v < min {
			min = v
		}
	}
	for i := 0; i < n; i++ {
		v := array[i]
		if v < max {
			r = i
		} else if v == max {
			if array[i-1] == v {

			} else {
				r = i - 1
			}
		}
		if v > max {
			max = v
		}
	}

	return []int{l, r}
}
