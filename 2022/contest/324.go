package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(smallestValue(15))
	//fmt.Println(smallestValue(12))
	//fmt.Println(smallestValue(11))
	//fmt.Println(smallestValue(3))
	//fmt.Println(smallestValue(2))
	//fmt.Println(smallestValue(100))
	//fmt.Println(smallestValue(99951))

	//d := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912}

	fmt.Println(536870912 * 2)

}

func cycleLengthQueries(n int, q [][]int) []int {
	m := len(q)
	ans := make([]int, m)
	//d := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912}

	cal := func(v int) []int {
		r := []int{}
		for x := v; x >= 1; x /= 2 {
			r = append(r, x)
		}
		for i := 0; i < len(r)/2; i++ {
			r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
		}
		return r
	}

	for i, v := range q {
		a := cal(v[0])
		b := cal(v[1])
		j := -1
		for ; j < len(a)-1 && j < len(b)-1; j++ {
			if a[j+1] != b[j+1] {
				break
			}
		}
		ans[i] = len(a) - j + len(b) - j - 1
	}
	return ans
}

func isPossible(n int, e [][]int) bool {
	nx := make([]map[int]bool, n)
	for _, v := range e {
		if nx[v[0]] == nil {
			nx[v[0]] = map[int]bool{}
		}
		if nx[v[1]] == nil {
			nx[v[1]] = map[int]bool{}
		}
		nx[v[0]][v[1]] = true
		nx[v[1]][v[0]] = true
	}

	p := []int{}
	for i, v := range nx {
		if len(v)%2 == 1 {
			p = append(p, i)
		}
	}

	if len(p) == 0 {
		return true
	} else if len(p) == 2 {
		a, b := p[0], p[1]
		if !nx[a][b] {
			return true
		}
		for i := 0; i < n; i++ {
			if i == a || i == b {
				continue
			}
			if !nx[a][i] && !nx[b][i] {
				return true
			}
		}
	} else if len(p) == 4 {
		a, b, c, d := p[0], p[1], p[2], p[3]
		if !nx[a][b] && !nx[c][d] {
			return true
		}
		if !nx[a][c] && !nx[b][d] {
			return true
		}
		if !nx[a][d] && !nx[b][c] {
			return true
		}
	}

	return false
}

func smallestValue(n int) int {
	ans := n
	a := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
		53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		101, 103, 107, 109, 113, 127, 131, 137, 139, 149,
		151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199,
		211, 223, 227, 229, 233, 239, 241,
		251, 257, 263, 269, 271, 277, 281, 283, 293,
		307, 311, 313, 317, 331, 337, 347, 349,
		353, 359, 367, 373, 379, 383, 389, 397,
		401, 409, 419, 421, 431, 433, 439, 443, 449,
		457, 461, 463, 467, 479, 487, 491, 499,
		503, 509, 521, 523, 541, 547,
		557, 563, 569, 571, 577, 587, 593, 599,
		601, 607, 613, 617, 619, 631, 641, 643, 647,
		653, 659, 661, 673, 677, 683, 691,
		701, 709, 719, 727, 733, 739, 743,
		751, 757, 761, 769, 773, 787, 797,
		809, 811, 821, 823, 827, 829, 839,
		853, 857, 859, 863, 877, 881, 883, 887,
		907, 911, 919, 929, 937, 941, 947,
		953, 967, 971, 977, 983, 991, 997}

	t := n
	for {
		fmt.Println(t)
		nt := 0

	loop:
		for x := t; x > 1; {

			for {
				for _, d := range a {
					if x%d == 0 {
						x /= d
						nt += d
						continue loop
					}
				}
				nt += x
				x = 1
				continue loop
			}

		}
		if nt < ans {
			ans = nt
		}
		if nt == t {
			break
		}
		t = nt
	}

	return ans
}

func similarPairs(a []string) int {
	ans := 0
	d := map[string]int{}
	for _, v := range a {
		c := map[byte]bool{}
		for _, x := range v {
			c[byte(x)] = true
		}
		b := []byte{}
		for x, _ := range c {
			b = append(b, x)
		}
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		ans += d[string(b)]
		d[string(b)]++
	}
	return ans
}
