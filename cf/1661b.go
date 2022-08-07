package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1661(os.Stdin, os.Stdout)
}

func b1661(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x int
	mod := 32768

	dp := make([]int, mod+1)
	st := []int{32768}
	dp[32768] = 0
	for len(st) > 0 {
		ns := []int{}
		for _, v := range st {
			pre := v - 1
			if pre > 0 && dp[pre] == 0 {
				dp[pre] = dp[v] + 1
				ns = append(ns, pre)
			}

			if v%2 == 0 {
				d := v/2
				if dp[d] == 0 {
					dp[d] = dp[v]+1
					ns = append(ns,d)
				}

				div := ((v + mod) / 2) % mod
				if dp[div] == 0 {
					dp[div] = dp[v] + 1
					ns = append(ns, div)
				}
			}
		}
		st = ns
	}
	dp[0] = 0
	//Println(dp[mod/2])
	//Println(dp[mod/4])

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)

		Fprintln(out, dp[x])
	}
}
