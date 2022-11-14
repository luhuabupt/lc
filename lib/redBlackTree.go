package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	a := []int{5, 3, 1, 7, 4, 6, 10, 4}
	rbt := redblacktree.NewWithIntComparator()
	for i, v := range a {
		r, has := rbt.Ceiling(v)
		rv := -1
		if has {
			rv = r.Value.(int)
		}

		l, has := rbt.Floor(v)
		lv := -1
		if has {
			lv = r.Value.(int)
		}

		fmt.Println(i, v, lv, l, rv, r)

		rbt.Put(v, i)
	}
}
