package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1702(os.Stdin, os.Stdout)
}

func c1702(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var k int
	var x int
	var y int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &k)

		st, que := make([]int, n), make([][]int, k)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			st[i] = x
		}
		for i := 0; i < k; i++ {
			Fscan(in, &x, &y)
			que[i] = []int{x, y}
		}

		m := map[int][]int{}
		for i, v := range st {
			if len(m[v]) == 0 {
				m[v] = []int{i, i}
			} else {
				m[v][1] = i
			}
		}

		for _, v := range que {
			x, y := v[0], v[1]
			ans := "NO"
			if len(m[x]) > 0 && len(m[y]) > 0 {
				if m[x][0] < m[y][1] {
					ans = "YES"
				}
			}
			Fprintln(out, ans)
		}
	}
}
