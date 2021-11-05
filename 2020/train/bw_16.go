package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803))
}

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)

	sum := 0
	i, n := 0, len(arr)
	min, ans := 100001, 0

	for x := 0; x < 100001 && i < n; x++ {
		for ; i < n && arr[i] < x; i++ {
			sum += arr[i]
		}
		if abs(target, sum+x*(n-i)) < min {
			min, ans = abs(target, sum+x*(n-i)), x
		}
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func pathsWithMaxScore(board []string) []int {
	n := len(board)
	M := int(1e9 + 7)
	unit := [][]int{{1, 0}, {0, 1}, {1, 1}}

	dp := map[int]map[int][2]int{}
	for i := 0; i <= n; i++ {
		dp[i] = map[int][2]int{}
		for j := 0; j <= n; j++ {
			dp[i][j] = [2]int{}
		}
	}

	for loop := n - 1; loop >= 0; loop-- {
		for _, u := range [][]int{{-1, 0}, {0, -1}} {
			for i, j := loop, loop; i >= 0 && j >= 0; {
				if i == n-1 && j == n-1 {
					dp[i][j] = [2]int{0, 1}
				} else if board[i][j] != 'X' {
					cnt, mx := 0, 0
					for _, v := range unit {
						if i+v[0] < n && j+v[1] < n && dp[i+v[0]][j+v[1]][0] > mx {
							mx = dp[i+v[0]][j+v[1]][0]
						}
					}
					for _, v := range unit {
						if i+v[0] < n && j+v[1] < n && dp[i+v[0]][j+v[1]][0] == mx {
							cnt += dp[i+v[0]][j+v[1]][1]
						}
					}
					if cnt > 0 && board[i][j] <= '9' && board[i][j] >= '0' {
						mx += int(board[i][j] - '0')
					}

					dp[i][j] = [2]int{mx, cnt % M}
				}
				i += u[0]
				j += u[1]
			}
		}
	}

	return []int{dp[0][0][0], dp[0][0][1]}
}
