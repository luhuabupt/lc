package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1676(os.Stdin, os.Stdout)
}

func b1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		mi := 1 << 30
		row := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			if x < mi {
				mi = x
			}
			row = append(row, x)
		}
		for _, v := range row {
			ans += v - mi
		}

		Fprintln(out, ans)
	}
}
