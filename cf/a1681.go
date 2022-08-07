package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	a1681(os.Stdin, os.Stdout)
}

func a1681(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var m int
	var x int

	chk := func(a, b []int) bool {
		return a[len(a)-1] >= b[len(b)-1]
		//fa := false
		////for i := 0; i < len(a) && i < len(b); i++ {
		////	if a[len(a)-1-i] > b[len(b)-1-i] {
		////		fa = true
		////		break
		////	} else if a[len(a)-1-i] < b[len(b)-1-i] {
		////		break
		////	}
		////}
		//return fa
	}

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, b := []int{}, []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a = append(a, x)
		}
		Fscan(in, &m)
		for i := 0; i < m; i++ {
			Fscan(in, &x)
			b = append(b, x)
		}
		sort.Ints(a)
		sort.Ints(b)

		if chk(a, b) {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
		if !chk(b, a) {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}
