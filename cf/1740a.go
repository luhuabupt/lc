package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1740(os.Stdin, os.Stdout)
}

func a1740(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var a string
	var b string

	c := map[byte]int{'S': 0, 'M': 1, 'L': 2}

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a)
		Fscan(in, &b)

		ans := ""
		if a == b {
			ans = "="
		} else {
			e, f := a[len(a)-1], b[len(b)-1]
			if e == f {
				if e == 'S' {
					if len(a) > len(b) {
						ans = "<"
					} else {
						ans = ">"
					}
				} else {
					if len(a) < len(b) {
						ans = "<"
					} else {
						ans = ">"
					}
				}
			} else if c[e] > c[f] {
				ans = ">"
			} else {
				ans = "<"
			}
		}

		Fprintln(out, ans)
	}
}
