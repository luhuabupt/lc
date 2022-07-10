package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func main() {
	c1700(os.Stdin, os.Stdout)
}

func c1700(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int64

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a := make([]int64, n)
		mi := int64(math.MaxInt64)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
			if a[i] < mi {
				mi = a[i]
			}
		}
		for i := 0; i < n; i++ {
			a[i] -= mi
		}

		ans := mi
		if mi < 0 {
			ans = -mi
		}
		do := int64(0)
		for i := 1; i < n-1; i++ {
			if a[i] >= a[i-1]+a[i+1] && a[i]-a[i-1]+a[i+1] > do {
				do = a[i] - a[i-1] + a[i+1]
			}
		}
		ans += do
		for i := 0; i < n; i++ {
			a[i] += do
		}
		x := int64(0)
		for _, v := range a {
			if v > x {
				x = v
			}
		}
		ans += x

		//for i, vv := range a {
		//	if i == 0 {
		//		if vv > 0 {
		//			ans += vv
		//		} else if vv < 0 {
		//			ans -= vv
		//			t -= vv
		//		}
		//		continue
		//	}
		//	v := vv + t
		//	if v < 0 {
		//		ans -= v * 2
		//		t -= v
		//	} else if v > 0 {
		//		if i == n-1 {
		//			ans += v
		//			continue
		//		} else {
		//			if vv <= a[i-1]+a[i+1] {
		//				continue
		//			}
		//			do := vv - a[i-1] + a[i+1]
		//			ans += do
		//		}
		//	}
		//}

		Fprintln(out, ans)
	}
}
