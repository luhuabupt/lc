package main

import (
	"fmt"
	"strings"
)

func capitalizeTitle_(title string) string {
	arr := strings.Split(title, " ")
	ans := []byte{}
	for _, s := range arr {
		if len(s) <= 2 {
			for i := 0; i < len(s); i++ {
				x := s[i]
				if s[i] >= 'A' && s[i] <= 'Z' {
					x += 'a' - 'A'
				}
				ans = append(ans, x)
			}
		} else {
			x := s[0]
			if x >= 'A' && x <= 'Z' {
				x += 'a' - 'A'
			}
			ans = append(ans, x)
			for i := 1; i < len(s); i++ {
				x := s[i]
				if s[i] >= 'A' && s[i] <= 'Z' {
					x += 'a' - 'A'
				}
				ans = append(ans, x)
			}
		}
		ans = append(ans, ' ')
	}
	return string(ans)
}

func capitalizeTitle(title string) string {
	arr := strings.Split(title, " ")
	ans := []string{}
	for _, v := range arr {
		v = strings.ToLower(v)
		if len(v) >= 2 {
			v = strings.Title(v)
		}
		ans = append(ans, v)
	}
	return strings.Join(ans, " ")
}

func longestPalindrome(words []string) int {
	cnt := map[string]int{}
	for _, v := range words {
		cnt[v]++
	}
	ans, flag := 0, false
	for s, v := range cnt {
		rs := string([]byte{s[1], s[0]})
		if rs == s {
			ans += v / 2
			if v%2 == 1 {
				flag = true
			}
		} else {
			x := v
			if cnt[rs] < x {
				x = cnt[rs]
			}
			cnt[rs] = 0
			ans += x * 4
		}
	}
	if flag {
		ans += 2
	}
	return ans
}

func longestPalindrome_(words []string) int {
	cnt := map[string]int{}
	for _, v := range words {
		cnt[v]++
	}
	ans := 0
	flag := false
	for s, v := range cnt {
		rs := get(s)
		if rs == s {
			if v%2 == 0 {
				ans += v * 2
				cnt[s] = 0
			} else {
				ans += v*2 - 2
				cnt[s] = 1
				flag = true
			}
		} else {
			x := min(v, cnt[rs])
			cnt[s] -= x
			cnt[rs] -= x
			ans += x * 4
		}
	}
	if flag {
		ans += 2
	}
	return ans
}
func get(a string) string {
	ans := make([]byte, 2)
	ans[0] = a[1]
	ans[1] = a[0]
	return string(ans)
}

func main() {
	//fmt.Println(possibleToStamp([][]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 4, 3))
	//fmt.Println(possibleToStamp([][]int{{0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}}, 2, 8))
	fmt.Println(possibleToStamp([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		9, 4))

}

func possibleToStamp(grid [][]int, h int, w int) bool {
	if h == 1 && w == 1 {
		return true
	}
	m, n := len(grid), len(grid[0])
	cnt := map[int]map[int]int{}
	cnt[-1] = map[int]int{}
	for j := -1; j < n; j++ {
		cnt[-1][j] = 0
	}
	for i, arr := range grid {
		cnt[i] = map[int]int{-1: 0}
		for j, v := range arr {
			cnt[i][j] = cnt[i-1][j] + cnt[i][j-1] - cnt[i-1][j-1] + v
		}
	}
	//fmt.Println(cnt)
	if w > n || h > m {
		return cnt[m-1][n-1] == m*n
	}
	can := make([][]int, m)
	for i, arr := range grid {
		can[i] = make([]int, n)
		for j, v := range arr {
			if v == 0 {
				if i+h-1 < m && j+w-1 < n && cnt[i+h-1][j+w-1]-cnt[i+h-1][j-1]-cnt[i-1][j+w-1]+cnt[i-1][j-1] == 0 {
					can[i][j] = 1
					grid[i][j] = 2
				}
			}
		}
	}
	//fmt.Println(can)

	cntCan := map[int]map[int]int{}
	cntCan[-1] = map[int]int{}
	for j := -1; j < n; j++ {
		cntCan[-1][j] = 0
	}
	for i, arr := range can {
		cntCan[i] = map[int]int{-1: 0}
		for j, v := range arr {
			cntCan[i][j] = cntCan[i-1][j] + cntCan[i][j-1] - cntCan[i-1][j-1] + v
		}
	}
	//fmt.Println(cntCan)

	for i, arr := range grid {
		for j, v := range arr {
			if v == 0 {
				x := cntCan[i][j]
				if j-w >= -1 {
					x -= cntCan[i][j-w]
				}
				if i-h >= -1 {
					x -= cntCan[i-h][j]
					if j-w >= -1 {
						x += cntCan[i-h][j-w]
					}
				}
				if x == 0 {
					//fmt.Println(i, j,cntCan[i][j], cntCan[i-h][j])
					return false
				}
			}
		}
	}

	return true
}
