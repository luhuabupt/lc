package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	//testRm()
	testRm2()
}

func testRm2() {
	a := [][]int{{0}, {1, 2}}
	rbt := redblacktree.NewWith(func(ai, bi interface{}) int {
		a := ai.([]int)
		b := bi.([]int)
		if len(a) < len(b) || len(a) == len(b) && a[len(a)-1] < b[len(b)-1] {
			return 1
		}
		if len(a) > len(b) || len(a) == len(b) && a[len(a)-1] > b[len(b)-1] {
			return -1
		}
		return 0
	})

	rbt.Put(a[0], 0)
	rbt.Put(a[1], 1)
	fmt.Println(rbt.String())
	rbt.Remove(a[0])
	fmt.Println(rbt.String())
}

func testCf() {
	a := []int{5, 3, 1, 7, 4, 6, 10, 4}
	rbt := redblacktree.NewWithIntComparator()
	for i, v := range a {
		r, has := rbt.Ceiling(v)
		rv := -1
		if has {
			rv = r.Value.(int)
		}

		l, has2 := rbt.Floor(v)
		lv := -1
		if has2 {
			lv = l.Value.(int)
		}

		fmt.Println(i, v, lv, l, rv, r)

		rbt.Put(v, i)
	}
}

func testRm() {
	rbt := redblacktree.NewWithIntComparator()

	rbt.Put(0, 0)
	rbt.Put(3, 3)
	//rbt.Put(1, 1)
	//rbt.Put(2, 2)
	fmt.Println(rbt.String())
	rbt.Remove(3)
	fmt.Println(rbt.String())
}
