package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1736(os.Stdin, os.Stdout)
}

func c1736(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)

		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		ans := 0
		s := 0
		l := 0
		for e, v := range a {
			l++
			for s < e && v < l {
				s++
				l--
			}
			ans += e - s + 1
		}

		Fprintln(out, ans)
	}
}
