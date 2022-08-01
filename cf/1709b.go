package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1709(os.Stdin, os.Stdout)
}

func b1709(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var m int
	var x int
	var s int
	var e int

	Fscan(in, &n)
	Fscan(in, &m)

	a := []int{}
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		a = append(a, x)
	}
	dl := make([]int64, n)
	dr := make([]int64, n)
	sl, sr := int64(0), int64(0)
	for i, _ := range a {
		if i == 0 {
			continue
		}
		if a[i] < a[i-1] {
			sl += int64(a[i-1] - a[i])
		}
		dl[i] = sl
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			sr += int64(a[i+1] - a[i])
		}
		dr[i] = sr
	}

	for l := 0; l < m; l++ {
		Fscan(in, &s, &e)
		s--
		e--
		ans := dl[e] - dl[s]
		if e < s {
			ans = dr[e] - dr[s]
		}
		Fprintln(out, ans)
	}
}
