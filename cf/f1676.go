package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	f1676(os.Stdin, os.Stdout)
}

func f1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var k int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &k)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}
		sort.Ints(a)

		m := map[int]int{}
		for _, v := range a {
			m[v]++
		}

		do := [][]int{}
		pre, cnt := a[0], 1
		for i := 1; i < n; i++ {
			v := a[i]
			if v == pre {
				cnt++
			} else {
				do = append(do, []int{pre, cnt})
				pre = v
				cnt = 1
			}
		}
		do = append(do, []int{pre, cnt})

		ans := []int{0, 0}
		l := 0
		for _, v := range do {
			if v[1] >= k {
				if m[v[0]-1] < k || l == 0 {
					l = v[0]
					if ans[0] == 0 {
						ans = []int{l, l}
					}
				} else if l > 0 && v[0]-l > ans[1]-ans[0] {
					ans = []int{l, v[0]}
				}
			} else {
				l = 0
			}
		}

		if ans[0] == 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, ans[0], ans[1])
		}
	}
}
