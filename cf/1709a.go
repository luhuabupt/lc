package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1709(os.Stdin, os.Stdout)
}

func a1709(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var q int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &q)
		a := []int{}
		for i := 0; i < 3; i++ {
			Fscan(in, &x)
			a = append(a, x)
		}

		do := 1
		for s := q - 1; ; {
			s = a[s]
			if s == 0 {
				break
			} else {
				s -= 1
				do++
			}
		}

		ans := "NO"
		if do == 3 {
			ans = "YES"
		}
		Fprintln(out, ans)
	}
}
