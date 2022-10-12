package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	c1719(os.Stdin, os.Stdout)
}

func c1719(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int
	var x int
	var y int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &m)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		w := make([][]int, n)
		mi := 0
		for i := 1; i < n; i++ {
			if a[i] > a[mi] {
				mi = i
				w[i] = append(w[i], i)
			} else {
				w[mi] = append(w[mi], i)
			}
		}

		for i := 0; i < m; i++ {
			Fscan(in, &x)
			x--
			Fscan(in, &y)

			t := y
			if t > n-1 {
				t = n - 1
			}
			ans := sort.SearchInts(w[x], t+1)
			if y > n-1 && mi == x {
				ans += y - n + 1
			}
			Fprintln(out, ans)
		}
	}
}
