package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" ")
}

func makeSimilar(a []int, b []int) int64 {
	ar, br := [2][]int{}, [2][]int{}
	for _, v := range a {
		ar[v%2] = append(ar[v%2], v)
	}
	for _, v := range b {
		br[v%2] = append(br[v%2], v)
	}
	do := func(a, b []int) int {
		r := 0
		sort.Ints(a)
		sort.Ints(b)
		for i, v := range a {
			if v < b[i] {
				r += b[i] - v
			}
		}
		return r
	}

	ans := do(ar[0], ar[0])
	ans += do(ar[1], br[1])
	return int64(ans / 2)
}

func makeSimilar_(a []int, b []int) int64 {
	//abs := func(a int) int {
	//	if a > 0 {
	//		return a
	//	}
	//	return -a
	//}
	oa, ob := 0, 0
	ar := []int{}
	br := []int{}
	ae := []int{}
	be := []int{}
	for _, v := range a {
		if v%2 == 1 {
			oa += v
			ar = append(ar, v)
		} else {
			ae = append(ae, v)
		}
	}
	for _, v := range b {
		if v%2 == 1 {
			ob += v
			br = append(br, v)
		} else {
			be = append(be, v)
		}
	}

	ans := 0
	if oa != ob {
		ans += abs(oa - ob)
	}

	do := func(a, b []int) int {
		sort.Ints(a)
		sort.Ints(b)
		t := 0
		ans := 0
		for i, v := range a {
			if v < b[i] { // +
				// -
				if t < 0 {
					t -= b[i] - v
					ans += b[i] - v
				} else {
					if t >= b[i]-v {
						t -= b[i] - v
					} else {
						ans += b[i] - v - t
						t -= b[i] - v
					}
				}
			}
			//else if v > b[i] { // -
			//	// +
			//	if t > 0 {
			//		t += v - b[i]
			//		ans += v - b[i]
			//	} else {
			//		if t+v-b[i] < 0 {
			//			t += v - b[i]
			//		} else {
			//			ans += t + v - b[i]
			//			t += v - b[i]
			//		}
			//	}
			//}
		}
		return ans
	}

	ans += do(ar, br)
	ans += do(ae, be)
	return int64(ans)

	//sort.Ints(a)
	//sort.Ints(b)
	//n := len(a)
	//da := make([]int, n)
	//db := make([]int, n)
	//for i, v := range a {
	//	if i == 0 {
	//		da[i] = v
	//	} else {
	//		da[i] = da[i-1] + v
	//	}
	//}
	//for i, v := range b {
	//	if i == 0 {
	//		db[i] = v
	//	} else {
	//		db[i] = db[i-1] + v
	//	}
	//}
	//
	//var re func(i, j int) int
	//re = func(i, j int) int {
	//	if i > j {
	//		return 0
	//	}
	//	sa := da[j]
	//	sb := db[j]
	//	if i > 0 {
	//		sa -= da[i]
	//		sb -= db[i]
	//	}
	//	if sa < sb {
	//
	//	}
	//
	//	m := i + j
	//	m /= 2
	//
	//	return re(i, m) + re(m+1, j)
	//}
	//
	//ans := re(0, len(a))
	//return int64(ans)
	//
	//fmt.Println(a, b)
	//
	//t := 0
	//ans := 0
	//for i, v := range a {
	//	if v < b[i] { // +
	//		// -
	//		if t < 0 {
	//			t -= b[i] - v
	//			ans += b[i] - v
	//		} else {
	//			if t >= b[i]-v {
	//				t -= b[i] - v
	//			} else {
	//				ans += b[i] - v - t
	//				t -= b[i] - v
	//			}
	//		}
	//	} else if v > b[i] { // -
	//		// +
	//		if t > 0 {
	//			t += v - b[i]
	//			ans += v - b[i]
	//		} else {
	//			if t+v-b[i] < 0 {
	//				t += v - b[i]
	//			} else {
	//				ans += t + v - b[i]
	//				t += v - b[i]
	//			}
	//		}
	//	}
	//	fmt.Println(i, v, t, ans)
	//}
	//
	//return int64(ans / 2)
}

func minCost(ar []int, c []int) int64 {
	n := len(ar)
	a := make([][2]int, n)

	for i, v := range ar {
		a[i] = [2]int{v, c[i]}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	dl := make([]int, n)
	dr := make([]int, n)
	l := make([]int, n)
	r := make([]int, n)

	l[0] = a[0][1]
	for i := 1; i < n; i++ {
		dl[i] += dl[i-1] + (a[i][0]-a[0][0])*a[i][1]
		l[i] = l[i-1] + a[i][1]
	}
	for i := n - 2; i >= 0; i-- {
		dr[i] = dr[i+1] + (a[n-1][0]-a[i][0])*a[i][1]
		r[i] = r[i+1] + a[i][1]
	}

	ans := 0
	for i, v := range a {
		t := dl[n-1] - dl[i] - (l[n-1]-l[i])*(v[0]-a[0][0])
		t += dr[0] - dr[i] - (r[0]-r[i])*(a[n-1][0]-v[0])

		if t < ans {
			ans = t
		}
	}

	return int64(ans)
}

func minCost_(ar []int, c []int) int64 {
	n := len(ar)
	a := make([][2]int, n)

	for i, v := range ar {
		a[i] = [2]int{c[i], v}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] > a[j][0]
	})

	ans := 0
	for i := 1; i < n; i++ {
		ans += abs(a[i][1]-a[0][1]) * a[i][0]
	}

	return int64(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func haveConflict(a []string, b []string) bool {
	st := func(v string) int {
		av := strings.Split(v, ":")
		h, _ := strconv.Atoi(av[0])
		m, _ := strconv.Atoi(av[1])
		return h*60 + m
	}

	for i := 0; i < 24*60; i++ {
		if i >= st(a[0]) && i <= st(a[1]) && i >= st(b[0]) && i <= st(b[1]) {
			return true
		}
	}

	return false
}
