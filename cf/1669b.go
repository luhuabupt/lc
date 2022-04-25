package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF16669B(os.Stdin, os.Stdout)
}

func CF16669B(_r io.Reader, _w io.Writer) {
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

		h := map[int]int{}
		for _, v := range r {
			h[v]++
			if h[v] >= 3 {
				Fprintln(out, v)
				continue loop
			}
		}

		Fprintln(out, -1)
	}
}
