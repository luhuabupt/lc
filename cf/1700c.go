package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	c1700(os.Stdin, os.Stdout)
}

func c1700(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int
	var x int

	for Fscan(in, &T); T > 0; T-- {

		Fscan(in, &n)
		Fscan(in, &m)

		a := make([]int, m)
		d := make([]int, n)
		for i := 0; i < m; i++ {
			Fscan(in, &x)
			a[i] = x
			d[x-1]++
		}
		sort.Ints(d)
		for i := 0; i < len(d)/2; i++ {
			d[i], d[len(d)-1-i] = d[len(d)-1-i], d[i]
		}

		ans := sort.Search(m*2, func(f int) bool {
			t := 0
			for _, v := range d {
				if v >= f {
					t += v - f
				} else {
					t -= (f - v) / 2
				}
				if t <= 0 {
					break
				}
			}
			return t <= 0
		})

		Fprintln(out, ans)
	}
}
