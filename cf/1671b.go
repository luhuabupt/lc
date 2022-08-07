package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF1671B(os.Stdin, os.Stdout)
}

func CF1671B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	var n int

loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		r := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r = append(r, x)
		}

		pre := r[0] + 1
		for i := 1; i < n; i++ {
			if r[i]-pre > 2 {
				Fprintln(out, "NO")
				continue loop
			}
			if r[i]-pre == 2 {
				pre = r[i] - 1
			} else if r[i]-pre == 1 {
				pre = r[i]
			} else if r[i] == pre {
				pre = r[i] + 1
			}
		}

		Fprintln(out, "YES")
	}
}
