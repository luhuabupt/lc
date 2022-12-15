package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1742(os.Stdin, os.Stdout)
}

func c1742(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string

loop:
	for Fscan(in, &T); T > 0; T-- {
		g := [8]string{}
		for i := 0; i < 8; i++ {
			Fscan(in, &s)
			g[i] = s
		}

		for i := 0; i < 8; i++ {
			if g[i] == "RRRRRRRR" {
				Fprintln(out, "R")
				continue loop
			}
		}

	lp:
		for j := 0; j < 8; j++ {
			for i := 0; i < 8; i++ {
				if g[i][j] != 'B' {
					continue lp
				}
			}
			Fprintln(out, "B")
			continue loop
		}

		Fprintln(out, "")
	}
}
