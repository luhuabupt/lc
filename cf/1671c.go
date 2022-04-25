package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	CF1671C(os.Stdin, os.Stdout)
}

func CF1671C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	var n int
	var t int64

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &t)

		r := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r = append(r, x)
		}
		sort.Ints(r)

		ans := int64(0)
		s := int64(0)
		for i, v := range r {
			s += int64(v)
			if t < s {
				break
			}
			ans += (t-s)/int64(i+1) + 1
		}

		Fprintln(out, ans)
	}
}
