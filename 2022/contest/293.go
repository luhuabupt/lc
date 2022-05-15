package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func removeAnagrams(w []string) []string {
	chk := func(a, b string) bool{
		c, d := [26]int{}, [26]int{}
		for _, v:= range a {
			c[v-'a']++
		}
		for _, v := range b {
			d[v-'a']++
		}
		for i := 0; i < 26;i++ {
			if c[i]!=d[i] {
				return false
			}
		}
		return true
	}

	for {
		nw := []string{w[0]}
		for i := 1; i < len(w); i++ {
			if chk(w[i], w[i-1]) {

			} else {
				nw = append(nw, w[i])
			}
		}
		if len(nw) == len(w) {
			break
		}
		w = nw
	}
	return w
}


func maxConsecutive(b int, t int, s []int) int {
	sort.Ints(s)
	ans := s[0]-b
	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] -1> ans {
			ans = s[i]-s[i-1]-1
		}
	}
	if t - s[len(s)-1] >ans {
		ans = t - s[len(s)-1]
	}
	return ans
}

func largestCombination(c []int) int {
	d := [31]int{}
	for i := 0;i < 31; i++ {
		for _, v := range c {
			if v &(1<<i) > 0 {
				d[i]++
			}
		}
	}
	ans := 0
	for _, v := range d {
		if v > ans {
			ans = v
		}
	}
	return ans
}



type CountIntervals struct {
	tr *redblacktree.Tree
	s int
}


func Constructor() CountIntervals {
	return CountIntervals{redblacktree.NewWithIntComparator(), 0}
}


func (t *CountIntervals) Add(left int, right int)  {

}


func (t *CountIntervals) Count() int {

}


/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */