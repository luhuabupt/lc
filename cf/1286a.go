package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	a1286(os.Stdin, os.Stdout)
}

func a1286(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x int
	Fscan(in, &n)

	a := make([]int, n)
	o, e := n/2+(n%2), n/2
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		a[i] = x
		if x%2 == 1 {
			o--
		} else if x > 0 {
			e--
		}
	}
	if n == 1 {
		Fprintln(out, 0)
		return
	}
	if o+e == n {
		Fprintln(out, 1)
		return
	}

	ans := 0
	b := [2][]int{}
	p, pl := -1, 0
	for i := 0; i < n; i++ {
		if a[i] == 0 { // 0
			pl++
			continue
		} else if i == 0 {
			p = a[i] % 2
			continue
		}

		// a[i] > 0
		if a[i-1] > 0 { // diff
			if a[i]%2 != a[i-1]%2 {
				ans++
			}
		} else { // 0 v
			if p == -1 {
				p = a[i] % 2
			}
			if p == a[i]%2 {
				b[p] = append(b[p], pl)
			} else {
				ans++
			}
			pl = 0
		}

		p = a[i] % 2
	}

	if pl > 0 {
		b[p] = append(b[p], pl)
	}

	sort.Ints(b[0])
	sort.Ints(b[1])

	Fprintln(out, o, e, b)

	for i, v := range b[0] {
		if v <= e {
			e -= v
		} else {
			ans += (len(b[0]) - i) * 2
			break
		}
	}
	for i, v := range b[1] {
		if v <= o {
			o -= v
		} else {
			ans += (len(b[1]) - i) * 2
			break
		}
	}

	Fprintln(out, ans)
}
