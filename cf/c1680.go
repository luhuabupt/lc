package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1680(os.Stdin, os.Stdout)
}

func c1680(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var sa string

	for Fscan(in, &T); T > 0; T-- {

		Fscan(in, &sa)
		s := []byte(sa)

		n := len(s)
		i := 0
		for ; i < n && s[i] == '0'; i++ {
		}
		if i == n {
			Fprintln(out, 0)
			continue
		} else {
			s = s[i:]
		}
		for i = len(s) - 1; i >= 0 && s[i] == '0'; i-- {
		}
		s = s[:i+1]

		n = len(s)
		d := make([]int, n)
		z := make([]int, n)
		c := 0
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				c++
			}
			d[i] = c
			z[i] = i-c+1
		}

		ans := 0
		for _, v := range s {
			if v == '0' {
				ans++
			}
		}

		w := 1
		pre := 0
		for i := 1; i < n; i++ {
			Println(i, s[i], pre, w, ans)
			if s[i] == '1' {
				// do rm 1, left 0
				x := max(c-d[i]+d[pre], z[i]-z[pre])
				ans = min(ans, x)
			}
			if s[i] == '1' {
				w++
			} else {
				w--
			}
			if s[i] == '0' && w <= 0 {
				pre = i
			}
		}

		Fprintln(out, ans)
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
