package main

import "fmt"

type tree struct {
	arr []int
}

func (t tree) add(i, v int) {
	fmt.Println(i)
	for i < len(t.arr) {
		t.arr[i] = t.arr[i] + v
		fmt.Println(lowbit(i), i+lowbit(i))
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
	x := []int{0, 2, 1, 4, 3}
	t := tree{make([]int, len(x)+1)}

	for _, v := range x {
		v = v + 1
		t.add(v, 1)
		//fmt.Println(v, t.query(1, v-1))
	}
	fmt.Println(t.arr)
}
