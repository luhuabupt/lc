package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a377(os.Stdin, os.Stdout)
}

func a377(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var m int
	var k int
	var x string

	Fscan(in, &m)
	Fscan(in, &n)
	Fscan(in, &k)

	g := make([][]byte, m)
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
		Fscan(in, &x)
		g[i] = append(g[i], []byte(x)...)
	}

	p := [][2]int{}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		vis[i][j] = true
		for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				//Fprintln(out, "next ", x, y, g[x][y])
				if !vis[x][y] && g[x][y] == '.' {
					dfs(x, y)
				}
			}
		}
		p = append(p, [2]int{i, j})
	}

	for i, ar := range g {
		for j, v := range ar {
			if v == '.' {
				dfs(i, j)
				//Fprintln(out, p)
				for t := 0; t < k; t++ {
					g[p[t][0]][p[t][1]] = 'X'
				}
				for _, a := range g {
					Fprintln(out, string(a))
				}
				return
			}
		}
	}
}
