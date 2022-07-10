package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1699(os.Stdin, os.Stdout)
}

func a1699(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)

		if x&1 == 1 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, 0, 0^(x/2), 0^(x/2))
		}
	}
}
