package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1680(os.Stdin, os.Stdout)
}

func b1680(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var m int
	var n int
	var s string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &m)
		Fscan(in, &n)

		g := make([]string, m)
		for i := 0; i < m; i++ {
			Fscan(in, &s)
			g[i] = s
		}

		ma := 1 << 30

	loop:
		for _, v := range g {
			for i, x := range v {
				if x == 'R' {
					ma = i
					break loop
				}
			}
		}

		ans := "YES"

	lp2:
		for j := 0; j < n; j++ {
			for i := 0; i < m; i++ {
				if g[i][j] == 'R' {
					if j < ma {
						ans = "NO"
						break lp2
					}
				}
			}
		}

		Fprintln(out, ans)
	}
}
