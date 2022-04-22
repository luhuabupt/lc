package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF16669C(os.Stdin, os.Stdout)
}

func CF16669C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	var n int

loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		r := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r = append(r, x)
		}

		for i := 0; i < len(r); i += 2 {
			if r[i]%2 != r[0]%2 {
				Fprintln(out, "NO")
				continue loop
			}
		}
		for i := 1; i < len(r); i += 2 {
			if r[i]%2 != r[1]%2 {
				Fprintln(out, "NO")
				continue loop
			}
		}

		Fprintln(out, "YES")
	}
}
