package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1676(os.Stdin, os.Stdout)
}

func a1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()


	var T int
	var x string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		a,b := 0, 0
		for i := 0;i <3; i++ {
			a += int(x[i]-'0')
		}
		for i := 3;i <6; i++ {
			b += int(x[i]-'0')
		}

		ans := "NO"
		if a == b {
			ans = "YES"
		}
		Fprintln(out, ans)
	}
}
