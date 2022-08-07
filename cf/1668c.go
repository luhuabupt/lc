package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1668(os.Stdin, os.Stdout)
}

func c1668(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var x, n int64
	Fscan(in, &n)

	row := []int64{}
	for i := int64(0); i < n; i++ {
		Fscan(in, &x)
		row = append(row, x)
	}

	ans := int64(1<<63 - 1)
	for z := int64(0); z < n; z++ {
		pre, t := int64(0), int64(0)
		for i := z - 1; i >= 0; i-- {
			t += pre/row[i] + 1
			pre = (pre/row[i] + 1) * row[i]
		}
		pre = 0
		for i := z + 1; i < n; i++ {
			t += pre/row[i] + 1
			pre = (pre/row[i] + 1) * row[i]
		}
		if t < ans {
			ans = t
		}
	}

	Fprintln(out, ans)
}
