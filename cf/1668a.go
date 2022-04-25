package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

//6
//1 1
//2 1
//1 3
//4 2
//4 6
//10 5

func main() {
	a1668(os.Stdin, os.Stdout)
}

func a1668(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x, y int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)

		if x < y {
			x, y = y, x
		}
		if y == 1 && x > 2 {
			Fprintln(out, -1)
		} else if x-y <= 1 {
			Fprintln(out, x+y-2)
		} else {
			a := 2*y - 1
			lt := x - y - 1
			if lt%2 == 0 {
				a += 4 * lt / 2
			} else {
				a += 4*(lt/2) + 3
			}

			Fprintln(out, a)
		}
	}
}
