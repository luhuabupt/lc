package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(robotWithString("zza"))
	fmt.Println(robotWithString("bac"))
	fmt.Println(robotWithString("bdda"))
	fmt.Println(robotWithString("vzhofnpo")) //"fnohopzv"
}

func robotWithString(s string) string {
	c := [26]int{}
	for _, v := range s {
		c[v-'a']++
	}

	st := []byte{}
	//ct := [26]int{}

	ans := []byte{}
loop:
	for _, v := range s {
		st = append(st, byte(v))
		//ct[v-'a']++
		c[v-'a']--

		for len(st) > 0 {
			cur := st[len(st)-1]
			st = st[:len(st)-1]

			for j := 'a'; j < int32(cur); j++ {
				if c[j-'a'] > 0 {
					st = append(st, cur)
					continue loop
				}
			}

			ans = append(ans, cur)
		}
	}
	return string(ans)
}

func robotWithString_(s string) string {
	d := [26][]int{}
	for i, v := range s {
		d[v-'a'] = append(d[v-'a'], i)
	}

	ans := []byte{}
	st := []byte{}

loop:
	for i, v := range s {
		st = append(st, byte(v))
		for len(st) > 0 {
			cur := st[len(st)-1]
			st = st[:len(st)-1]

			//fmt.Println("   ", i, cur, st)
			flag := false
			// Y add st, continue
			// Eq add ans, continue
			// N add ans, st pop
			for j := 0; j < int(cur-'a'); j++ {
				x := sort.SearchInts(d[j], i+1)
				if x < len(d[j]) {
					flag = true
					st = append(st, cur)
					continue loop
				}
			}

			// N
			if !flag {
				ans = append(ans, cur)
			}

			// eq
			x := sort.SearchInts(d[cur-'a'], i+1)
			if x < len(d[cur-'a']) {
				//fmt.Println(" eq loop ", cur, x, len(d[cur-'a']))
				continue loop
			}
		}
	}

	return string(ans)
}

func numberOfPaths(g [][]int, k int) int {
	m, n := len(g), len(g[0])
	M := int(1e9 + 7)
	dp := make([][][]int, m)
	for i, ar := range g {
		dp[i] = make([][]int, n)
		for j, v := range ar {
			dp[i][j] = make([]int, k)
			if i-1 > 0 { // shang
				for x := 0; x < k; x++ {
					dp[i][j][(x+v)%k] = dp[i-1][j][x]
					dp[i][j][(x+v)%k] %= M
				}
			}
			if j-1 > 0 { // zuo
				for x := 0; x < k; x++ {
					dp[i][j][(x+v)%k] = dp[i-1][j-1][x]
					dp[i][j][(x+v)%k] %= M
				}
			}
		}
	}
	return dp[m-1][n-1][0]
}
