package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1719(os.Stdin, os.Stdout)
}

func b1719(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var k int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &k)

		if k%4 == 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			if k%4 == 2 {
				for i := 0; i*4 < n; i++ {
					a := i*4 + 1
					b, c, d := a+1, a+2, a+3
					Fprintln(out, b, a)
					if c < n {
						Fprintln(out, c, d)
					}
				}
			} else {
				for i := 0; i*4 < n; i++ {
					a := i*4 + 1
					b, c, d := a+1, a+2, a+3
					Fprintln(out, a, b)
					if c < n {
						Fprintln(out, c, d)
					}
				}
			}
		}
	}
}
