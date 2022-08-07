package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1706(os.Stdin, os.Stdout)
}

func a1706(_r io.Reader, _w io.Writer) {
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
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
		}
		ans := make([]byte, m)
		for i := 0; i < m; i++ {
			ans[i] = 'B'
		}
		for _, v := range r {
			a, b := v-1, m-v
			if a > b {
				a, b = b, a
			}
			if ans[a] == 'A' {
				ans[b] = 'A'
			} else {
				ans[a] = 'A'
			}
		}

		Fprintln(out, string(ans))
	}
}
