package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1743(os.Stdin, os.Stdout)
}

func b1743(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a := []int{}
		for i := 1; i <= n; i += 2 {
			a = append(a, i)
		}
		for i := n / 2 * 2; i >= 2; i -= 2 {
			a = append(a, i)
		}
		for _, v := range a {
			Fprintf(out, "%d ", v)
		}
		Fprintln(out, "")
	}
}
