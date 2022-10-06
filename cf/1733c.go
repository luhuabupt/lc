package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1733(os.Stdin, os.Stdout)
}

func c1733(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		r := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
		}

		st := []int{}
		for i, v := range r {
			if len(st) == 0 || v > r[st[len(st)-1]] {
				st = append(st, i)
			}
		}

		ans := [][]int{}
		si := len(st) - 1

		Fprintln(out, r, st)

		for i := n - 1; i >= 0; i-- {
			if st[si] < i {
				si--
			}
			if st[si] != i {
				ans = append(ans, []int{st[si], i})
			}
		}

		Fprintln(out, len(ans))
		if len(ans) > 0 {
			for _, v := range ans {
				Fprintln(out, v[0], v[1])
			}
		}
	}
}
