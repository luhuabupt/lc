package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF16669A(os.Stdin, os.Stdout)
}

func CF16669A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		o := "Division "
		if x < 1400 {
			o += "4"
		} else if x < 1600 {
			o += "3"
		} else if x < 1900 {
			o += "2"
		} else {
			o += "1"
		}

		Fprintln(out, o)

	}
}
