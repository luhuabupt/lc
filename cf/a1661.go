package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1661(os.Stdin, os.Stdout)
}

func a1661(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(v int) int64 {
		if v < 0 {
			return int64(-v)
		}
		return int64(v)
	}

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			b[i] = x
		}

		ans := int64(1<<63 - 1)

		num := n
		for kth := 0; kth <= n/2; kth++ {
			mask := (1 << kth) - 1
			limit := 1 << num
			for mask < limit {
				//Printf("%b ", mask)
				ta := append([]int{}, a...)
				tb := append([]int{}, b...)
				for j := 0; j < n; j++ {
					if mask>>j&1 == 1 {
						ta[j], tb[j] = tb[j], ta[j]
					}
				}
				//Println(i, ta, tb)
				t := int64(0)
				for j := 1; j < n; j++ {
					t += abs(ta[j] - ta[j-1])
					t += abs(tb[j] - tb[j-1])
				}
				if t < ans {
					ans = t
				}


				if mask == 0 {
					break
				}
				lowBit := mask & -mask
				right := mask + lowBit
				mask = (((right ^ mask) >> 2) / lowBit) | right
			}
		}

		Fprintln(out, ans)
	}
}
