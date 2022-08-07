package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1670(os.Stdin, os.Stdout)
}

func b1670(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var s string
	var cn int
	var x string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &s)
		Fscan(in, &cn)
		b := [26]bool{}
		for i:= 0; i < cn; i++ {
			Fscan(in, &x)
			Println(s, x, x[0]-'a')
			b[x[0]-'a'] = true
		}

		sa := []byte(s)
		ans := 0
		for len(sa) > 0{
			if b[sa[0]-'a'] {
				break
			}
			ns := []byte{}
			for i := 0; i < len(sa)-1; i++ {
				if b[sa[i+1]-'a'] {
					ns = append(ns, sa[i])
				}
			}
			ns = append(ns, sa[len(sa)-1])
			sa = ns
			ans++
		}

		Fprintln(out, ans)
	}
}
