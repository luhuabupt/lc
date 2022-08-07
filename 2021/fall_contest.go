package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"....X.", "....X.", "XOOO..", "......", "......"}
	a = []string{".X.", ".O.", "XO."}
	fmt.Println(flipChess(a))
}

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	m1, m2 := map[int]int{}, map[int]int{}

	for _, arr := range source {
		for _, v := range arr {
			m1[v]++
		}
	}
	for _, arr := range target {
		for _, v := range arr {
			m2[v]++
		}
	}

	common := 0
	for k, v := range m1 {
		common += min(v, m2[k])
	}

	return len(source)*len(source[0]) - common
}

func maxmiumScore(cards []int, cnt int) int {
	sort.Ints(cards)
	n := len(cards)

	ans := 0
	for i := 0; i < cnt; i++ {
		ans += cards[n-1-i]
	}

	if ans%2 == 0 {
		return ans
	}

	// 找个偶数替换
	sc1 := 0
	for i := n - cnt; i < n; i++ {
		if cards[i]%2 == 0 {
			for j := n - 1 - cnt; j >= 0; j-- {
				if cards[j]%2 == 1 {
					sc1 = ans + cards[j] - cards[i]
					break
				}
			}
			break
		}
	}

	// 找个奇数替换
	sc2 := 0
	for i := n - cnt; i < n; i++ {
		if cards[i]%2 == 1 {
			for j := n - 1 - cnt; j >= 0; j-- {
				if cards[j]%2 == 0 {
					sc2 = ans + cards[j] - cards[i]
					break
				}
			}
			break
		}
	}

	if sc2 > sc1 {
		return sc2
	}
	return sc1
}

func flipChess(chessboard []string) int {
	m, n := len(chessboard), len(chessboard[0])
	cb := make([][]byte, m)
	for i, arr := range chessboard {
		cb[i] = make([]byte, n)
		for j, _ := range arr {
			cb[i][j] = arr[j]
		}
	}

	al := []int{0}
	for i, arr := range cb {
		for j, v := range arr {
			if v == '.' {
				al = append(al, getChange(cb, i, j))
			}
		}
	}
	sort.Ints(al)
	return al[len(al)-1]
}

func getChange(cb [][]byte, i int, j int) int {
	xd := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	yd := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	m, n := len(cb), len(cb[0])

	cc := make([][]byte, m)
	for i, arr := range cb {
		cc[i] = make([]byte, n)
		for j, _ := range arr {
			cc[i][j] = arr[j]
		}
	}

	change := [][]int{[]int{i, j}}
	for len(change) > 0 {
		head := change[0]
		change = change[1:]
		for k := 0; k < 8; k++ {
			xi, yi := head[0]+xd[k], head[1]+yd[k]
			for xi >= 0 && xi < m && yi >= 0 && yi < n && cc[xi][yi] == 'O' {
				yi += yd[k]
				xi += xd[k]
			}
			if xi >= 0 && xi < m && yi >= 0 && yi < n && cc[xi][yi] == 'X' {
				for ci, cj := head[0]+xd[k], head[1]+yd[k]; cc[ci][cj] == 'O'; ci += xd[k] {
					cc[ci][cj] = 'X'
					change = append(change, []int{ci, cj})
					cj += yd[k]
				}
			}
		}
	}

	black := 0
	for i, arr := range cc {
		for j, v := range arr {
			if v == 'X' && cb[i][j] == 'O' {
				black++
			}
		}
	}
	return black
}

func circleGame(toys [][]int, circles [][]int, r int) int {
	m, ans := map[int]map[int]bool{}, 0
	for _, arr := range circles {
		if m[arr[0]] == nil {
			m[arr[0]] = map[int]bool{}
		}
		m[arr[0]][arr[1]] = true
	}
	for _, arr := range toys {
		if arr[2] > r {
			continue
		}
		flag := false
		for i := arr[0] + arr[2] - r; i <= arr[0]-arr[2]+r; i++ {
			for j := arr[1] + arr[2] - r; j <= arr[1]-arr[2]+r; j++ {
				if m[i][j] && checkD(i, j, arr[0], arr[1], r, arr[2]) {
					ans++
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
	}
	return ans
}
func checkD(i, j, a, b, r, x int) bool {
	if (i-a)*(i-a)+(j-b)*(j-b) <= (r-x)*(r-x) {
		return true
	}
	return false
}
