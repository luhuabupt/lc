package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1676(os.Stdin, os.Stdout)
}

func d1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &m)
		Fscan(in, &n)

		g := make([][]int, m)
		for i := 0; i < m; i++ {
			g[i] = make([]int, n)
			for j := 0; j < n; j++ {
				Fscan(in, &x)
				g[i][j] = x
			}
		}

		l := make([]int, m+n-1)
		r := make([]int, m+n-1)

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				l[i+j] += g[i][j]
				r[i-j+n-1] += g[i][j]
			}
		}

		ans := 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				t := l[i+j] + r[i-j+n-1] - g[i][j]
				if t > ans {
					ans = t
				}
			}
		}

		Fprintln(out, ans)
	}
}
