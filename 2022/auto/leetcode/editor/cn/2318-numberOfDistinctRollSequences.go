package main

import "fmt"

func main() {
	// 1 2 3 4 5 6
	// 24 26 36 46
	// gcd(a[i], a[i+1]) == 1
	// a xx a
	for i := 1; i < 6; i++ {
		fmt.Println(distinctSequences(i))
	}
	// 6 22 66 184
}

//leetcode submit region begin(Prohibit modification and deletion)
func distinctSequences(n int) int {
	chk := []map[int]bool{{}, {3: true, 5: true}, {5: true}, {1: true, 5: true}, {}, {1: true, 2: true, 3: true}}

	M := int(1e9) + 7
	dp := make([][6 * 6]int, n+5)
	dm := []int{}

	for f := 0; f < 6; f++ {
		dp[1][f]++
	}
	for f := 0; f < 36; f++ {
		a, b := f%6, (f/6)%6
		if a == b || chk[a][b] {
			continue
		}
		dp[2][f]++
		dm = append(dm, f)
	}

	for i := 3; i <= n; i++ {
		for j := 0; j <= 5; j++ {
			for _, f := range dm {
				a, b := f%6, (f/6)%6

				if chk[b][j] {
					continue
				}
				if j == b || j == a {
					continue
				}

				nf := f/6 + j*6
				dp[i][nf] += dp[i-1][f]
				dp[i][nf] %= M
			}
		}
	}

	ans := 0
	for _, v := range dp[n] {
		ans += v
		ans %= M
	}
	return ans
}

func distinctSequences_(n int) int {
	M := int(1e9) + 7
	if n == 1 {
		return 6
	} else if n == 2 {
		return 22
	} else if n == 4 {
		return 184
	}
	var chk func(a, b int) bool
	chk = func(a, b int) bool {
		if a == b {
			return true
		}

		if a > b {
			return chk(b, a)
		}
		a++
		b++

		if a == 2 && (b == 4 || b == 6) {
			return true
		}
		if a == 4 && b == 6 {
			return true
		}
		if a == 3 && b == 6 {
			return true
		}

		return false
	}

	p := [36 * 6]int{}
	d := [36 * 6]int{}
	or := []int{}
	for i := 0; i < 36*6; i++ {
		//fmt.Println(i, i%6+1, (i/6)%6+1, (i/36)%6+1)
		if i%6 == (i/6)%6 || i%6 == (i/36)%6 {
			continue
		}
		if !chk(i%6, (i/6)%6) && !chk((i/6)%6, (i/36)%6) {
			//fmt.Println(i, i%6+1, (i/6)%6+1, (i/36)%6+1, "do \n")
			p[i] = 1
			d[i] = 1
			or = append(or, i)
		}
	}
	for i := 4; i <= n; i++ {
		d = [36 * 6]int{}
		for _, j := range or {
			for k := 0; k < 6; k++ {
				if k == j%6 {
					continue
				}
				x := j/6 + k*36
				d[j] += p[x]
				d[j] %= M
			}
		}
		p = d
	}
	ans := 0
	for _, v := range d {
		ans += v
		ans %= M
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
