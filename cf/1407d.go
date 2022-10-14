package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1407(os.Stdin, os.Stdout)
}

func d1407(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x int
	Fscan(in, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		Fscan(in, &x)
		a[i] = x
	}

	idx := 0
	vis := make([]bool, n)
	st := []int{0}
	for {
		ns := []int{}
		for _, v := range st {
			if v == n-2 {
				Fprintln(out, idx)
				return
			}
			if !vis[v+1] {
				ns = append(ns, v+1)
				vis[v+1] = true
			}

			if a[v+1] == a[v] {
				continue
			}
			if a[v+1] > a[v] {
				for j := v + 1; j < n; j++ {
					if a[j] >= a[j-1] {

					} else {
						break
					}
				}
			}
		}
	}

}
