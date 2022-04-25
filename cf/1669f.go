package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF16669F(os.Stdin, os.Stdout)
}

func CF16669F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	var n int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		row := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			row = append(row, x)
		}

		ans := 0
		s := 0
		l, r := map[int]int{}, map[int]int{}
		for i, v := range row {
			s += v
			l[s] = i
		}
		s = 0
		for i := 0; i < n; i++ {
			v := row[n-1-i]
			s += v
			r[s] = i
		}

		for v, i := range l {
			if j, ok := r[v]; ok {
				if i+j+2 <= n {
					ans = i + j + 2
				}
			}
		}

		Fprintln(out, ans)
	}
}
