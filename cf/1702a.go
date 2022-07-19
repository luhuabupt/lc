package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1702(os.Stdin, os.Stdout)
}

func a1702(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		for i := 1; i <= int(1e9); i *= 10 {
			if x >= i && x < i*10 {
				Fprintln(out, x-i)
			}
		}
	}
}
