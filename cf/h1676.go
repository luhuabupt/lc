package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	h1676(os.Stdin, os.Stdout)
}

func h1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x-1
		}

		ans := 0
		for i, v := range a {
			for j := i +1; j < n;j++ {
				if a[j] < v {
					ans++
				}
			}
		}
		m := map[int]int{}
		for _,  v := range a {
			m[v]++
		}
		for _, v := range m {
			if v > 1 {
				ans +=  v * (v-1)/2
			}
		}

		Fprintln(out, ans)
	}
}
