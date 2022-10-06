package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1730(os.Stdin, os.Stdout)
}

func a1730(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var c int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &c)
		r := make([]int, 101)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[x]++
		}
		ans := 0
		for _, v := range r {
			if v < c {
				ans += v
			} else {
				ans += c
			}
		}

		Fprintln(out, ans)
	}
}
