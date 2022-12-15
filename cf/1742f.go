package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	f1742(os.Stdin, os.Stdout)
}

func f1742(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var s string
	var xi int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a, b := [26]int{}, [26]int{}
		a[0], b[0] = 1, 1
		sa := 1
		sb := 1

	loop:
		for i := 0; i < n; i++ {
			Fscan(in, &xi, &x)
			Fscan(in, &s)
			if xi == 1 {
				for _, v := range s {
					a[v-'a'] += x
					sa += x
				}
			}
			if xi == 2 {
				for _, v := range s {
					b[v-'a'] += x
					sb += x
				}
			}
			//Fprintln(out, a, b)

			c := 0
			for ; c < 26; c++ {
				if a[c] > 0 {
					// g
					for d := c + 1; d < 26; d++ {
						if b[d] > 0 {
							Fprintln(out, "YES")
							continue loop
						}
					}
					if sb == b[c] && b[c] == a[c] {
						Fprintln(out, "NO")
						continue loop
					}
					if a[c] == sa && a[c] <= b[c] {
						Fprintln(out, "YES")
						continue loop
					}
				}
			}

			Fprintln(out, "NO")
		}

	}
}
