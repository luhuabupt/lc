package main

import (
	"fmt"
)

type pair struct {
	w string
	c int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.c < b.c || a.c == b.c && a.w > b.w }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	a := "abce"
	fmt.Println(len(a))
	for i, v := range a {
		fmt.Println(v)
		fmt.Println(int32(v))
		fmt.Println(int(v))
		fmt.Println(a[i])
	}
}

func testVar() {
	a := []int{0, 1, 2, 3}
	for len(a) > 0 {
		var t int
		t, a = a[0], a[1:]
		t, a := a[0], a[1:]
		fmt.Println(t, a)
	}
}
