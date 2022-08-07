package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(appealSum("abbca"))
	fmt.Println(appealSum("code"))
}

func appealSum(s string) int64 {
	n := len(s)
	m := [26][]int{}
	for i, v := range s {
		m[v-'a'] = append(m[v-'a'], i)
	}

	dp := make([]int, n)
	sa := make([]int, n)
	dp[0] = 1
	sa[0] = 1
	for i, v := range s {
		if i == 0 {
			continue
		}
		x := v - 'a'
		ar := m[x]
		y := sort.SearchInts(ar, i)

		dp[i] = dp[i-1]
		if y-1 >= 0 {
			dp[i] += i - ar[y-1]
		} else {
			dp[i] += i + 1
		}
		sa[i] = sa[i-1] + dp[i]
	}
	return int64(sa[n-1])
}

func countDistinct(a []int, k int, p int) int {
	n := len(a)
	s := make([]int, n)
	m := map[string]bool{}

	sa := make([]string, n)
	for i, v := range a {
		sa[i] = strconv.Itoa(v)
	}

	if a[0]%p == 0 {
		s[0] = 1
	}
	for i := 1; i < n; i++ {
		s[i] = s[i-1]
		if a[i]%p == 0 {
			s[i]++
		}
	}

	for l := 0; l < n; l++ {
		for i := 0; i < n-l; i++ {
			x := s[i+l]
			if i > 0 {
				x -= s[i-1]
			}
			if x <= k {
				mk := strings.Join(sa[i:i+l+1], "_")
				m[mk] = true
			}
		}
	}

	return len(m)
}

func removeDigit(a string, d byte) string {
	n := len(a)
	ans := ""
	for i := 0; i < n; i++ {
		if a[i] == d {
			if a[:i]+a[i+1:] > ans {
				ans = a[:i] + a[i+1:]
			}
		}
	}
	return ans
}

func minimumCardPickup(a []int) int {
	n := len(a)
	m := map[int][]int{}
	for i, v := range a {
		m[v] = append(m[v], i)
	}
	if len(m) == n {
		return -1
	}
	ans := n
	for i, v := range a {
		x := sort.SearchInts(m[v], i)
		if x+1 < len(m[v]) {
			if m[v][x+1]-i < ans {
				ans = m[v][x+1] - i
			}
		}
	}
	return ans + 1
}
