package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	c1682(os.Stdin, os.Stdout)
}

func c1682(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		row := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			row[i] = x
		}
		sort.Ints(row)
		a, b := []int{row[n-1]}, []int{row[n-1]}
		for i := n - 2; i >= 0; i-- {
			for ; i >= 0 && row[i] == a[len(a)-1]; i-- {
				continue
			}
			if i < 0 {
				break
			}
			a = append(a, row[i])
			i--

			for ; i >= 0 && row[i] == b[len(b)-1]; i-- {
				continue
			}
			if i < 0 {
				break
			}
			b = append(b, row[i])
		}

		Fprintln(out, min(len(a), len(b)))
	}
}
