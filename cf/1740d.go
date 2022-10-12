package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1740(os.Stdin, os.Stdout)
}

func d1740(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	diff := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		r := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
		}

		ans := 0
		for l := 1; l <= n; l *= 2 {
			for i := 0; i+l < n; i += l * 2 {
				if diff(r[i], r[i+l]) >= 2*l {
					//Fprintln(out, r, l, i)
					Fprintln(out, -1)
					continue loop
				}
				if r[i] > r[i+l] {
					ans++
				}
			}
		}

		Fprintln(out, ans)
	}
}
