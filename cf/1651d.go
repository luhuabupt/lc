package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1651(os.Stdin, os.Stdout)
}

// 0 1 2 3 4 3 2
// 1 1 2 2 1 0 2 6

func d1651(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n)

	st := make([][2]int, n)
	m := map[int]map[int][3]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &x, &y)
		st[i] = [2]int{x, y}
		if _, ok := m[x]; !ok {
			m[x] = map[int][3]int{}
		}
		m[x][y] = [3]int{0, 0, 0}
	}

	dir := [][]int{{0, 1}, {0, -1}, {-1, 0}, {0, -1}}
	deep := 0
	for len(st) > 0 {
		for _, v := range st {
			i, j := v[0], v[1]
			for _, d := range dir {
				x, y := i+d[0], j+d[1]
				if m[x] == nil {
					m[i][j] = [3]int{deep, x, y}
				}
				if _, ok := m[x][y]; !ok {
					m[i][j] = [3]int{deep, x, y}
				}
			}
		}
	}

	for _, v := range ans {
		Fprintf(out, "%d ", v)
	}
	Fprintln(out, "")
}
