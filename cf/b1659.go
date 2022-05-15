package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1659(os.Stdin, os.Stdout)
}

func b1659(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var k int
	var x string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &x)
		ans := []byte{}
		se := make([]int, n)
		if k%2 == 0 {
			for i, v := range x {
				if v == '0' {
					if k == 0 {
						ans = append(ans, x[i:]...)
						break
					}
					if k > 0 {
						k--
						se[i]++
					}
				}
				ans = append(ans, '1')
			}
		} else {
			for i, v := range x {
				if k == 0 {
					if v == '1' {
						ans = append(ans, '0')
					} else {
						ans = append(ans, '1')
					}
				}
				if k > 0 {
					if v == '1' {
						se[i]++
						k--
					}
					ans = append(ans, '1')
				}
			}
		}

		se[0] += k - k%2
		if k > 0 && k%2 == 1 {
			se[n-1]++
			ans[n-1] = '0'
		}

		Fprintln(out, string(ans))
		for i, v := range se {
			if i > 0 {
				Fprintf(out, " ")
			}
			Fprintf(out, "%d", v)
		}
		Fprintln(out, "")
	}
}
