package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1733(os.Stdin, os.Stdout)
}

func b1733(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int
	var y int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &x)
		Fscan(in, &y)

		if x == 0 && y == 0 || x*y != 0 {
			Fprintln(out, -1)
		} else {
			if x == 0 {
				x = y
			}
			if (n-1)%x != 0 {
				Fprintln(out, -1)
			} else {
				p := 0
				for i := 1; i < n; i++ {
					if i > 0 {
						Fprintf(out, " ")
					}
					if i%x == 1 {
						p = i
					}
					Fprintf(out, "%d", p+1)
				}
				Fprintln(out, "")
			}
		}
	}
}
