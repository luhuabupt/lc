package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	e1676(os.Stdin, os.Stdout)
}

func e1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &m)

		r := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
		}
		sort.Ints(r)

		a := make([]int, n)
		sum :=0
		for i := n - 1; i >= 0; i-- {
			sum += r[i]
			a[n-1-i] = sum
		}

		for j := 0; j < m; j++ {
			Fscan(in, &x)
			ans := sort.SearchInts(a, x)
			if ans == n {
				Fprintln(out, -1)
			} else {
				Fprintln(out, ans+1)
			}
		}
	}
}
