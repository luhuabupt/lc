package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1700(os.Stdin, os.Stdout)
}

func a1700(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &m)
		Fscan(in, &n)
		ans := int64(0)
		for i := 0; i < n; i++ {
			ans += int64(i + 1)
		}
		for j := 2; j <= m; j++ {
			ans += int64(j * n)
		}

		Fprintln(out, ans)
	}
}
