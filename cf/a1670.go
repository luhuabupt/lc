package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1670(os.Stdin, os.Stdout)
}

func a1670(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		neg := 0
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			if x < 0 {
				neg++
				x = -x
			}
			a[i] = x
		}

		for i := 1; i < neg; i++ {
			if a[i] > a[i-1] {
				Fprintln(out, "NO")
				continue loop
			}
		}
		for i := neg+1; i < n;  i++ {
			if a[i] < a[i-1] {
				Fprintln(out, "NO")
				continue loop
			}
		}

		Fprintln(out, "YES")
	}
}
