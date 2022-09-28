package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	e1619(os.Stdin, os.Stdout)
}

// 0 1 2 3 4 3 2
// 1 1 2 2 1 0 2 6

func e1619(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		c := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			c[x]++
		}

		st := []int{}
		ans := make([]int, n+1)
		do := 0
		for i := 0; i <= n; i++ {
			// cal
			ans[i] = do + c[i]

			// construct
			if i < n && c[i] == 0 {
				if len(st) == 0 {
					for j := i + 1; j <= n; j++ {
						ans[j] = -1
					}
					break
				}
				x := st[len(st)-1]
				st = st[:len(st)-1]
				do += i - x
				continue
			}
			for j := 1; j < c[i]; j++ {
				st = append(st, i)
			}
		}

		for _, v := range ans {
			Fprintf(out, "%d ", v)
		}
		Fprintln(out, "")
	}
}
