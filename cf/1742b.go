package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1742(os.Stdin, os.Stdout)
}

func b1742(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		r := map[int]bool{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[x] = true
		}
		ans := "NO"
		if len(r) == n {
			ans = "YES"
		}
		Fprintln(out, ans)
	}
}
