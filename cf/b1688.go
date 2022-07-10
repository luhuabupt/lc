package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	b1688(os.Stdin, os.Stdout)
}

func b1688(_r io.Reader, _w io.Writer) {
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
		oc := 0
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
			if x%2 == 0 {
				oc++
			}
		}
		if oc == 0 {
			Fprintln(out, 0)
			continue loop
		}
		if oc < n {
			Fprintln(out, oc)
			continue loop
		}

		sort.Ints(a)
		mi := 32
		for _, v := range a {
			t := 0
			for v%2 == 0 {
				v /= 2
				t++
				if t > mi {
					break
				}
			}
			if t < mi {
				mi = t
			}
		}

		Fprintln(out, oc+mi-1)
	}
}
