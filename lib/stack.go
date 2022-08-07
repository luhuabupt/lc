package main

import "fmt"

func main() {
	fmt.Println(bigSL([]int{9, 8, 3, 5, 7}))
	fmt.Println(bigSR([]int{9, 8, 3, 5, 7}))
}

// 左侧严格大于
func bigSL(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && a[st[len(st)-1]] <= a[i] {
			st = st[:len(st)-1]
		}
		ans[i] = -1
		if len(st) > 0 {
			ans[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}

// 右侧严格大于
func bigSR(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	st := []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && a[st[len(st)-1]] <= a[i] {
			st = st[:len(st)-1]
		}
		ans[i] = -1
		if len(st) > 0 {
			ans[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
