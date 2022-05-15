package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1674(os.Stdin, os.Stdout)
}

func a1674(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	var y int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		Fscan(in, &y)

		a, b := 0, 0
		if x > y || y%x != 0 {
			Fprintln(out, a, b)
		} else {
			Fprintln(out, 1, y/x)
		}

	}
}
