package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1700(os.Stdin, os.Stdout)
}

func a1700(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		a := []int{}
		s := 0
		for i := 0; i < 4; i++ {
			Fscan(in, &x)
			s += x
			a = append(a, x)
		}

		ans := 0
		if s == 4 {
			ans = 2
		} else if s != 0 {
			ans = 1
		}

		Fprintln(out, ans)
	}
}
