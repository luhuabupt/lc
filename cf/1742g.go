package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	g1742(os.Stdin, os.Stdout)
}

func g1742(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		r := make([]int, n)
		ans := make([]int, n)
		mi := 0
		ua := map[int]bool{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
			if x > r[mi] {
				mi = i
			}
		}
		ans[0] = r[mi]
		ua[mi] = true

		mx := r[mi]
		mxi := 0
		for l := 1; l < 31; l++ {
			t := mx
			for i, v := range r {
				if v|mx > t {
					t = v | mx
					mxi = i
				}
			}
			if t == mx {
				break
			}
			mx = t
			ua[mxi] = true
			ans[l] = r[mxi]
		}

		idx := len(ua)
		for i, v := range r {
			if ua[i] {
				continue
			}
			ans[idx] = v
			idx++
		}

		for _, v := range ans {
			Fprintf(out, "%d ", v)
		}
		Fprintln(out, "")
	}
}
