package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1743(os.Stdin, os.Stdout)
}

func c1743(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T int
	var n int
	var s string
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &s)

		r := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r = append(r, x)
		}

		a, b := 0, 0
		if s[0] == '1' {
			a = r[0]
		}

		for i := 1; i < n; i++ {
			if s[i] == '0' {
				continue
			}
			tb := b
			if s[i-1] == '0' {
				b = max(a, b) + r[i-1]
			} else {
				b = b + r[i-1]
			}
			a = max(a, tb) + r[i]
			//Fprintln(out, "i ", i, a, b)
		}

		Fprintln(out, max(a, b))
	}
}
