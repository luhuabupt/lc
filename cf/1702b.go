package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1702(os.Stdin, os.Stdout)
}

func b1702(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		ans := 0
		a := map[int32]bool{}
		for _, v := range x {
			if a[v] {
				continue
			}
			if len(a) < 3 {
				a[v] = true
			} else {
				ans++
				a = map[int32]bool{v: true}
			}
		}
		if len(a) > 0 {
			ans++
		}
		Fprintln(out, ans)

	}
}
