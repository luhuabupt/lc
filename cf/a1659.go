package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1659(os.Stdin, os.Stdout)
}

func a1659(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int
	var y int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		ans := []byte{}
		t := x / (y + 1)
		flag := false
		if t*(y+1) != x {
			flag = true
			t++
		}
		do := y - (x - (x/(y+1))*(y+1))
		for ; y > 0; y-- {
			if flag && y == do {
				t--
			}
			for i := 0; i < t && x > 0; i++ {
				ans = append(ans, 'R')
				x--
			}
			ans = append(ans, 'B')
		}
		for ; x > 0; x-- {
			ans = append(ans, 'R')
		}

		Fprintln(out, string(ans))
	}
}
