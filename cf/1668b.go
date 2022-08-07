package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	b1668(os.Stdin, os.Stdout)
}

func b1668(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x, n int
	var m int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)

		row := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			row = append(row, x)
		}
		sort.Ints(row)

		u := int64(1 + 2*row[n-1])
		for i := n - 2; i >= 0; i-- {
			v := row[i]
			u += int64(1 + v)
		}
		u -= int64(row[0])

		if u > m {
			Fprintln(out, "NO")

		} else {
			Fprintln(out, "YES")
		}
	}
}
