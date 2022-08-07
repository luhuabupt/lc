package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1706(os.Stdin, os.Stdout)
}

func d1706(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var k int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &k)
		ar := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			ar[i] = x
		}
	}
}
