package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1682(os.Stdin, os.Stdout)
}

func a1682(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var s string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &s)
		mi := n / 2
		ans := 0
		for i := mi; i < n; i++ {
			if s[mi] == s[i] {
				ans++
			} else {
				break
			}
		}
		for i := mi - 1; i >= 0; i-- {
			if s[mi] == s[i] {
				ans++
			} else {
				break
			}
		}

		Fprintln(out, ans)
	}
}
