package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1699(os.Stdin, os.Stdout)
}

func b1699(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var m int
	var n int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &m)
		Fscan(in, &n)

		ans := make([][]int, m)
		for i := 0; i < m; i++ {
			ans[i] = make([]int, n)
		}

	}
}
