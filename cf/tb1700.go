package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	tb1700(os.Stdin, os.Stdout)
}

func tb1700(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int64

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a := make([]int64, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		r := int64(0)
		for i := 1; i < n; i++ {
			v := r + a[i]
			if v > a[i-1] {
				r += v - a[i-1]
			}
		}

	}
}
