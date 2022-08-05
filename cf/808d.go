package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d808(os.Stdin, os.Stdout)
}

func d808(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var x int64
	Fscan(in, &n)
	rs := make([]int64, n)
	tt := int64(0)
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		rs[i] = x
		tt += x
	}

	ans := "NO"
	m := map[int64]bool{0: true}
	sm := int64(0)
	for _, v := range rs {
		sm += v
		m[v] = true
		left := tt - sm
		if sm > left && (sm-left)%2 == 0 && m[(sm-left)/2] {
			ans = "YES"
			break
		}
	}

	if ans == "YES" {
		Fprintln(out, ans)
		return
	}

	m = map[int64]bool{0: true}
	sm = 0
	for i := n - 1; i >= 0; i-- {
		sm += rs[i]
		m[rs[i]] = true
		left := tt - sm
		if sm > left && (sm-left)%2 == 0 && m[(sm-left)/2] {
			ans = "YES"
			break
		}
	}

	Fprintln(out, ans)
}
