package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1694(os.Stdin, os.Stdout)
}

func a1694(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	var y int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		Fscan(in, &y)

		ans := []byte{}
		for x > 0 && y > 0 {
			ans = append(ans, '0')
			ans = append(ans, '1')
			x--
			y--
		}
		for x > 0 {
			ans = append(ans, '0')
			x--
		}
		for y > 0 {
			ans = append(ans, '1')
			y--
		}

		Fprintln(out, string(ans))
	}
}
