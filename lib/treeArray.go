package main

import "fmt"

type tree struct {
	arr []int
}

func (t tree) add(i, v int) {
	for i < len(t.arr) {
		t.arr[i] = t.arr[i] + v
		i += lowbit(i)
	}
}

func (t tree) sum(i int) (ans int) {
	for i > 0 {
		ans += t.arr[i]
		i -= lowbit(i)
	}
	return
}

func (t tree) query(l, r int) (ans int) {
	return t.sum(r) - t.sum(l-1)
}

func lowbit(v int) int {
	return v & -v
}

func main() {
	x := []int{1, 2, 3, 4, 5, 1}
	t := tree{make([]int, len(x)+1)}

	for i, v := range x {
		t.add(i+1, i+v+1000)
	}
	fmt.Println(t.arr)
	fmt.Println(t.query(1, 2))
	fmt.Println(t.query(2, 3))
	fmt.Println(t.query(1, 5))
}
