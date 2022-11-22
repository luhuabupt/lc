package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1743(os.Stdin, os.Stdout)
}

func a1743(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		ans := (10 - n) * (10 - n - 1) / 2 * 6

		Fprintln(out, ans)
	}
}
