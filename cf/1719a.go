package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1719(os.Stdin, os.Stdout)
}

func a1719(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	var y int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		Fscan(in, &y)

		ans := "Burenka"
		if (x+y)%2 == 0 {
			ans = "Tonya"
		}

		Fprintln(out, ans)
	}
}
