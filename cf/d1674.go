package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	d1674(os.Stdin, os.Stdout)
}

func d1674(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		row := make([]int, n)
		st := make([]int, n)

		for i := 0; i < n; i++ {
			Fscan(in, &x)
			row[i] = x
			st[i] = x
		}
		sort.Ints(st)

		chk := func(i, j int) bool {
			return st[i] == row[i] && st[j] == row[j] || st[i] == row[j] && st[j] == row[i]
		}
		if n%2 == 1 {
			if st[0] != row[0] {
				Fprintln(out, "NO")
				continue loop
			}
		}
		for i := n % 2; i < n; i += 2 {
			if !chk(i, i+1) {
				Fprintln(out, "NO")
				continue loop
			}
		}

		Fprintln(out, "YES")
	}
}
