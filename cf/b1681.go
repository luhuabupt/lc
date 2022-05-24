package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1681(os.Stdin, os.Stdout)
}

func b1681(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, b := []int{}, []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a = append(a, x)
		}
		Fscan(in, &m)
		for i := 0; i < m; i++ {
			Fscan(in, &x)
			b = append(b, x)
		}

		s := 0
		for _, v := range b {
			s += v
			s %= n
		}

		Fprintln(out, a[s])

	}
}
