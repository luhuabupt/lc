package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b280(os.Stdin, os.Stdout)
}

func b280(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var x int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		a[i] = x
	}

	l := bigL(a)
	r := bigR(a)

	ans := 0
	for i, v := range a {
		if l[i] > 0 && a[l[i]]^v > ans {
			ans = a[l[i]] ^ v
		}
		if r[i] > 0 && a[r[i]]^v > ans {
			ans = a[r[i]] ^ v
		}
	}
	Fprintln(out, ans)
}

func bigL(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && a[st[len(st)-1]] <= a[i] {
			st = st[:len(st)-1]
		}
		ans[i] = -1
		if len(st) > 0 {
			ans[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}

func bigR(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	st := []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && a[st[len(st)-1]] <= a[i] {
			st = st[:len(st)-1]
		}
		ans[i] = -1
		if len(st) > 0 {
			ans[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
