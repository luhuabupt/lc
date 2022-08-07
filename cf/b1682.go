package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1682(os.Stdin, os.Stdout)
}

func b1682(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		ans := 0
		r := make([]int, n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = i
		}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
		}

		if r[0] != 0 {
			Fprintln(out, ans)
			continue
		}

	loop:
		for i := 1; (1 << i) < n; i++ {
			for j := 1 << (i - 1); j < 1<<i && j < n; j++ {
				Fprintln(out, j, ans)
				if a[j] != r[j] {
					if ans != 0 {
						if ans != 1<<(i-1) {
							ans = 0
							break loop
						}
					} else {
						ans = 1 << (i - 1)
					}
				}
			}
		}

		Fprintln(out, ans)
	}
}
