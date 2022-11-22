package main

import "fmt"

func main() {
	fmt.Println(" ")
}

func countSubarrays_(a []int, c int, d int) int64 {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	if c > d {
		return 0
	}
	//if c == d {
	//	ans := 0
	//	t := 0
	//	for _, v := range a {
	//		if v == c {
	//			t++
	//		} else {
	//			t = 0
	//		}
	//		ans += t
	//	}
	//	ans += t
	//	return int64(ans)
	//}
	ar := [][]int{}
	cur := []int{}
	for _, v := range a {
		if v >= c && v <= d {
			cur = append(cur, v)
		} else {
			if len(cur) > 0 {
				ar = append(ar, cur)
			}
			cur = []int{}
		}
	}
	if len(cur) > 0 {
		ar = append(ar, cur)
	}

	ans := 0
	for _, arr := range ar {
		cm := []int{-1, -1}
		for i, v := range arr {
			if v == c {
				ans += cm[1] + 1
				cm[0] = i
			} else if v == d {
				ans += cm[0] + 1
				cm[1] = i
			} else {
				ans += min(cm[1], cm[0]) + 1
			}
		}
	}

	return int64(ans)
}

func sumOfNumberAndReverse(n int) bool {
	for i := 0; i < n; i++ {
		av := []int{}
		for v := i; v > 0; v /= 10 {
			av = append(av, v%10)
		}
		x := 0
		for j := 0; j < len(av); j++ {
			x *= 10
			x += av[j]
		}
		if x+i == n {
			return true
		}
	}
	return false
}

func countDistinctIntegers(a []int) int {
	n := len(a)
	for i := 0; i < n; i++ {
		v := a[i]
		av := []int{}
		for ; v > 0; v /= 10 {
			av = append(av, v%10)
		}
		j := 0
		for j < len(av) && av[j] == 0 {
			j++
		}
		x := 0
		for ; j < len(av); j++ {
			x *= 10
			x += av[j]
		}

		a = append(a, x)
	}

	c := map[int]bool{}
	for _, v := range a {
		c[v] = true
	}
	return len(c)
}
