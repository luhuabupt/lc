package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d75(os.Stdin, os.Stdout)
}

func d75(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, x int
	Fscan(in, &n, &m)
	rs := make([][]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &l)
		for j := 0; j < n; j++ {
			Fscan(in, &x)
			rs[i] = append(rs[i], x)
		}
	}

	arr := make([]int, m)
	for i := 0; i < m; i++ {
		Fscan(in, &x)
		arr[i] = x
	}

}
