package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	e1742(os.Stdin, os.Stdout)
}

func e1742(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &m)
		r := make([]int, n)
		q := make([]int, m)
		st := []int{}
		stv := []int{}
		s := make([]int64, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
			if len(st) == 0 || stv[len(stv)-1] < x {
				st = append(st, i)
				stv = append(stv, x)
			}

			s[i] += int64(x)
			if i > 0 {
				s[i] += s[i-1]
			}
		}
		for i := 0; i < m; i++ {
			Fscan(in, &x)
			q[i] = x
		}

		ans := make([]int64, m)
		for i, v := range q {
			t := sort.SearchInts(stv, v+1)
			if t >= len(stv) {
				ans[i] = s[n-1]
			} else if t == 0 {
				ans[i] = 0
			} else {
				ans[i] = s[st[t]-1]
			}
		}

		for _, v := range ans {
			Fprintf(out, "%d ", v)
		}
		Fprintln(out, "")
	}
}
