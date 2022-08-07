package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1709(os.Stdin, os.Stdout)
}

func c1709(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)

		n := len(s)
		sl, sr := 0, 0
		for _, v := range s {
			if v == '(' {
				sl++
			}
			if v == ')' {
				sr++
			}
		}

		ans := "YES"

		tl, tr := n/2-sl, n/2-sr
		if tl == 0 || tr == 0 {

		} else {
			l, r := 0, 0
			for i, v := range s {
				if v == '?' {
					if i == 0 || r >= l {
						// must left
						tl--
						if tl == 0 {
							break
						}

						l++
					} else {
						ans = "NO"
						break
					}
				}
				if v == '(' {
					l++
				}
				if v == ')' {
					r++
				}
			}
		}

		Fprintln(out, ans)
	}
}
