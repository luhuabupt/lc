package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	g1676(os.Stdin, os.Stdout)
}

func g1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int
	var y string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		row := make([]int, n)
		row[0] = 0
		for i := 1; i < n; i++ {
			Fscan(in, &x)
			row[i] = x-1
		}
		Fscan(in, &y)
		b := make([]bool, n)
		for i, v := range y {
			if v == 'B' {
				b[i] = true
			}
		}

		next := make([][]int, n)
		for i := 1; i < n; i++ {
			v := row[i]
			next[v] = append(next[v], i)
		}

		ans := 0
		var dfs func(i int) (int , int)
		dfs = func(i int) (int, int){
			ww, bb := 0, 0
			for _, v := range next[i] {
				vw, vb := dfs(v)
				ww += vw
				bb += vb
			}
			if b[i] {
				bb++
			} else {
				ww++
			}
			if ww == bb {
				ans++
			}
			return ww, bb
		}

		dfs(0)

		Fprintln(out, ans)
	}
}
