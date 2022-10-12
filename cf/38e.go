package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	e38(os.Stdin, os.Stdout)
}

func e38(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var x int

	Fscan(in, &n)
	a := make([]int, n)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		Fscan(in, &x)
		a[i] = x
		Fscan(in, &x)
		c[i] = x
	}

	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {

		}
	}

	Fprintln(out, ans)
}
