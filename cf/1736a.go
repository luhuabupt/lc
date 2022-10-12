package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1736(os.Stdin, os.Stdout)
}

func a1736(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		b := make([]int, n)
		sa := 0
		sb := 0
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
			sa += x
		}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			b[i] = x
			sb += x
		}
		if sa < sb {
			sa, sb = sb, sa
			a, b = b, a
		}

		ans := 0
		if sa == sb {
			for i := 0; i < n; i++ {
				if a[i] != b[i] {
					ans++
					break
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if a[i] != b[i] {
					a[i] = b[i]
					ans++

					if ans == sa-sb {
						break
					}
				}
			}
			for i := 0; i < n; i++ {
				if a[i] != b[i] {
					ans++
					break
				}
			}
		}

		Fprintln(out, ans)
	}
}
