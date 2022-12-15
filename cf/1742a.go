package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1742(os.Stdin, os.Stdout)
}

func a1742(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var a int
	var b int
	var c int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)

		if a == b+c || b == a+c || c == b+a {
			Fprintln(out, "YES")
			continue
		}

		Fprintln(out, "NO")
	}
}
