package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1688(os.Stdin, os.Stdout)
}

func a1688(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)

		ans := 0
		if x == 1 {
			ans = 3
		} else if x%2 == 0 {
			lb := x & -x
			ans = lb
			if lb == x {
				ans++
			}
		} else {
			ans = 1
		}

		Fprintln(out, ans)
	}
}
