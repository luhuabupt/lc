package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1680(os.Stdin, os.Stdout)
}

func a1680(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var l int
	var r int
	var a int
	var b int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l)
		Fscan(in, &r)
		Fscan(in, &a)
		Fscan(in, &b)

		ans := 0
		if l <= a {
			if a <= r {
				ans = a
			} else {
				ans = l + a
			}
		} else {
			if b < l {
				ans = a + l
			} else {
				ans = l
			}
		}

		Fprintln(out, ans)
	}
}
