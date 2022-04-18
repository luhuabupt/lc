package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(minimizeResult("247+38"))
	//fmt.Println(minimizeResult("12+34"))
	//fmt.Println(minimizeResult("999+999"))
	//fmt.Println(minimizeResult("123+991"))
	//
	//fmt.Println(maximumProduct([]int{0, 4}, 5))
	//fmt.Println(maximumProduct([]int{6, 2, 3, 2}, 4))
	//fmt.Println(maximumProduct([]int{6, 3, 3, 2}, 4))

	fmt.Println(maximumBeauty([]int{1, 1, 1, 3}, 7, 6, 12, 1))
	fmt.Println(maximumBeauty([]int{2, 3, 4, 5}, 10, 5, 2, 6))
	fmt.Println(maximumBeauty([]int{1, 5, 5, 9, 15}, 36, 12, 9, 2)) // 58
	fmt.Println(maximumBeauty([]int{2, 8}, 24, 18, 6, 3))           // 54
	fmt.Println(maximumBeauty([]int{13}, 18, 15, 9, 2))             // 54 // 30
	fmt.Println(maximumBeauty([]int{22323, 64818, 97718, 14354, 27837, 6347, 43299, 23010, 18590, 12706, 1579, 52401, 23473, 82978, 1012, 2248, 50247, 755, 12672, 57870, 90646, 87848, 71069, 82723, 83385, 66909, 39266, 97768, 62453, 30454, 68978, 88590, 11213, 74010, 65683, 75460, 3118, 98281, 28128, 84992, 52206, 92770, 74240, 33266, 41603, 19267, 36991, 86658, 43834, 84533, 75187, 31502, 61181, 31333, 37324, 13352, 51735, 89812, 56745, 44204, 85482, 70358, 48830, 91989, 82778, 34042, 3293, 63626, 41301, 43763, 39681, 91917, 40165, 57944, 34357, 22395, 26084, 21666, 40781, 28998, 12385, 10720, 66853, 42324, 28327, 30125, 15522, 12223}, 997843, 100000, 14880, 45790))
	//[5,5,15,1,9]
	//36
	//12
	//9
	//2
	//// 58
}

func minimizeResult(s string) string {
	as := strings.Split(s, "+")

	mi := 1 << 32
	ii, jj := 0, 0

	cal := func(i, j int) int {
		a, c := 1, 1
		if i > 0 {
			a, _ = strconv.Atoi(as[0][:i])
		}
		if j < len(as[1]) {
			c, _ = strconv.Atoi(as[1][j:])
		}
		x, _ := strconv.Atoi(as[0][i:])
		y, _ := strconv.Atoi(as[1][:j])
		return a * (x + y) * c
	}

	for i := 0; i < len(as[0]); i++ { // i 括号位置
		for j := 1; j <= len(as[1]); j++ {
			x := cal(i, j)
			if x < mi {
				mi = x
				ii, jj = i, j
			}
		}
	}
	return s[:ii] + "(" + as[0][ii:] + "+" + as[1][:jj] + ")" + as[1][jj:]
}

func maximumProduct(a []int, k int) int {
	sort.Ints(a)
	mi := sort.Search(1e7, func(i int) bool {
		t := k
		for _, v := range a {
			if v < i {
				t -= (i - v)
			} else {
				break
			}
		}
		return t < 0
	}) - 1
	for i, v := range a {
		if v < mi {
			a[i] = mi
			k -= (mi - v)
		} else {
			break
		}
	}
	for i, _ := range a {
		if k > 0 {
			a[i]++
			k--
		} else {
			break
		}
	}

	//fmt.Println(a)
	M := int(1e9) + 7
	ans := 1
	for _, v := range a {
		ans *= v
		ans %= M
	}
	return ans
}

func maximumBeauty(f []int, newFlowers int64, target int, full int, p int) int64 {
	nf := int(newFlowers)
	sort.Ints(f)
	n := len(f)
	if f[0] >= target {
		return int64(n * full)
	}

	sum, ss := []int{}, 0
	for _, v := range f {
		ss += v
		sum = append(sum, ss)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ans := 0
	pu := nf
	for i := n - 1; i >= 0; i-- { // per
		if n == 0 {
			if min(f[0]+pu, target-1)*p > full {
				return int64(min(f[0]+pu, target-1) * p)
			}
		}
		if target > f[i] {
			pu -= target - f[i]
			if pu < 0 {
				break // 不够了
			}
		}
		mi := 0
		if i > 0 {
			mi = sort.Search(target, func(j int) bool {
				idx := sort.SearchInts(f[:i], j)
				//fmt.Println(j, idx)
				if idx == 0 {
					return false
				}
				if sum[idx-1]+pu < j*(min(idx, i)) {
					//fmt.Println("false")
					return true
				}
				return false
			}) - 1
		}

		//fmt.Println(i, mi, pu, (n-i)*full+mi*p)
		if (n-i)*full+mi*p > ans {
			ans = (n-i)*full + mi*p
		}
	}

	x := sort.Search(target, func(j int) bool {
		idx := sort.SearchInts(f, j)
		//fmt.Println(j, idx)
		if idx == 0 {
			return false
		}
		if sum[idx-1]+nf < j*idx {
			//fmt.Println("false")
			return true
		}
		return false
	}) - 1
	if f[0] < target && x*p > ans {
		ans = x * p
	}

	return int64(ans)
}
