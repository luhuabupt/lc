package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	b1730(os.Stdin, os.Stdout)
}

func b1730(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		p := make([][]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			p[i] = make([]int, 2)
			p[i][0] = x
		}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			p[i][1] = x
		}

		sort.Slice(p, func(i, j int) bool {
			return p[i][0] < p[j][0] || p[i][0] == p[j][0] && p[i][1] > p[j][1]
		})
		np := [][]int{}
		for _, v := range p {
			if len(np) == 0 || v[0] != np[len(np)-1][0] {
				np = append(np, v)
			}
		}
		p = np
		n = len(p)

		l, r := make([]int, n), make([]int, n)
		mx := -int(1e10)
		for i := 0; i < n; i++ {
			mx = max(p[i][1]-p[i][0], mx)
			l[i] = mx
		}
		mx = -int(1e10)
		for i := n - 1; i >= 0; i-- {
			mx = max(p[i][1]+p[i][0], mx)
			r[i] = mx
		}

		//Fprintln(out, p, l, r)

		mi := int(1e10)
		ans := float64(p[0][0])
		for i := 1; i < n; i++ {
			if l[i]+r[i] < mi {
				mi = l[i] + r[i]
				ans = (float64(r[i]) - float64(l[i])) / float64(2)
			}
		}

		//li, ri := 0, n-1
		//for i := 0; i < n; i++ {
		//	if i*2 < n {
		//		if 2*p[1][i]+n-i*2 > 2*p[1][li]+n-li*2 {
		//			li = i
		//		}
		//	}
		//	if i > n/2 {
		//		if 2*p[1][i]+i*2-n > 2*p[1][ri]+ri*2-n {
		//			ri = i
		//		}
		//	}
		//}
		//ans := (p[1][ri]-p[1][ri]+ri+li)/2

		Fprintln(out, ans)
	}
}
