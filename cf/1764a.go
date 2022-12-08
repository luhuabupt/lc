package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1764(os.Stdin, os.Stdout)
}

func a1764(_r io.Reader, _w io.Writer) {
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

		ma := -1
		al, ar := 0, 0
		l, r := 0, 0
		t := -1
		p := a[0]
		for i := 1; i < n; i++ {
			if p == a[i] {
				t++
			} else {
				if t < 0 {
					t = -1
					l = i
				} else {

				}
			}
			if t > ma {
				al, ar = l, r
			}
		}

		Fprintln(out, al, ar)
	}
}
