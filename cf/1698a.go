package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1698(os.Stdin, os.Stdout)
}

func a1698(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
		}
		Fprintln(out, x)
	}
}
