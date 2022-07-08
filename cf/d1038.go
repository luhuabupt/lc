package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1038(os.Stdin, os.Stdout)
}

func d1038(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	if n == 1 {
		Fprintln(out, 1)
		return
	} else if n == 2 {
		Fprintln(out, 3)
		return
	}

	M := 998244353
	odd, even := 3, 3
	for i := 3; i <= n; i++ {
		if i%2 == 0 {
			even *= (i + 1)
			odd *= (i + 1)
		} else {
			even, odd = even*(i+1)/2+odd*(i+1)/2, even*(i+1)/2
		}
		odd %= M
		even %= M
	}

	Fprintln(out, odd)
}
