package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1733(os.Stdin, os.Stdout)
}

func a1733(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var k int
	var x int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &k)

		r := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
		}

		ma := make([]int, k)
		for i := 0; i < n; i++ {
			ma[i%k] = max(ma[i%k], r[i])
		}

		ans := int64(0)
		for _, v := range ma {
			ans += int64(v)
		}

		Fprintln(out, ans)
	}
}
