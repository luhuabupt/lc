package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	c1681(os.Stdin, os.Stdout)
}

func c1681(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, b := []int{}, []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a = append(a, x)
		}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			b = append(b, x)
		}

		ha, hb := map[int]bool{}, map[int]bool{}
		for _, v := range a {
			ha[v] = true
		}
		for _, v := range b {
			hb[v] = true
		}

		ra, rb := map[int]int{}, map[int]int{}
		d := 0
		for i := 1; i <= n; i++ {
			if ha[i] {
				d++
				ra[i] = d
			}
		}
		d = 0
		for i := 1; i <= n; i++ {
			if hb[i] {
				d++
				rb[i] = d
			}
		}

		ar, br := []int{}, []int{}
		for _, v := range a {
			ar = append(ar, ra[v])
		}
		for _, v := range b {
			br = append(br, rb[v])
		}

		do := [][]int{}
		for i, v := range ar {
			do = append(do, []int{v, br[i], i})
		}

		sort.Slice(do, func(i, j int) bool {
			return do[i][0] < do[j][0] || do[i][0] == do[j][0] && do[i][1] < do[j][1]
		})

		for i := 1; i < n; i++ {
			if do[i][1] < do[i-1][1] {
				Fprintln(out, -1)
				continue loop
			}
		}
		//Fprintln(out, do)

		o := make([]int, n)
		p := make([]int, n)
		for i := 0; i < n; i++ {
			p[i] = i
			o[i] = i
		}

		t := [][]int{}
		for i, v := range do {
			if o[i] != v[2] {
				j := p[v[2]]
				p[o[i]], p[v[2]] = j, i
				o[j], o[i] = o[i], o[j]
				t = append(t, []int{i, j})
			}
			//Fprintln(out, i, o)
		}

		Fprintln(out, len(t))
		for _, v := range t {
			Fprintln(out, v[0]+1, v[1]+1)
		}

		//// chk
		//Fprintln(out, a)
		//Fprintln(out, b)
		//for _, v := range t {
		//	a[v[0]], a[v[1]] = a[v[1]], a[v[0]]
		//	b[v[0]], b[v[1]] = b[v[1]], b[v[0]]
		//}
		//Fprintln(out, a)
		//Fprintln(out, b)
	}
}
