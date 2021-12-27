package main

func minimumBuckets(street string) int {
	n := len(street)
	ans := 0
	arr := []byte(street)
	for i := 0; i < n; i++ {
		if arr[i] == 'H' {
			if i == 0 {
				if i+1 >= n || arr[i+1] == 'H' {
					return -1
				}
				arr[i+1] = 'B'
			} else if i == n-1 {
				if arr[i-1] == 'H' {
					return -1
				} else {
					arr[i-1] = 'B'
				}
			} else {
				if arr[i+1] == 'H' && arr[i-1] == 'H' {
					return -1
				}
				if arr[i+1] == 'B' || arr[i-1] == 'B' {
					continue
				}
				if arr[i+1] == '.' {
					arr[i+1] = 'B'
				} else {
					arr[i-1] = 'B'
				}
			}
		}
	}
	for _, v := range arr {
		if v == 'B' {
			ans++
		}
	}
	return ans
}

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	ans := 0
	x1, y1, x2, y2 := startPos[0], startPos[1], homePos[0], homePos[1]
	for x1 < x2 {
		x1++
		ans += rowCosts[x1]
	}
	for x1 > x2 {
		x1--
		ans += rowCosts[x1]
	}
	for y1 < y2 {
		y1++
		ans += colCosts[y1]
	}
	for y1 > y2 {
		y1--
		ans += colCosts[y1]
	}
	return ans
}

func countPyramids(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	s := make([][]int, m)

	for i, arr := range grid {
		s[i] = make([]int, n)
		for j, v := range arr {
			s[i][j] = v
			if j > 0 {
				s[i][j] += s[i][j-1]
			}
		}
	}

	ans := 0
	for i := 0; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for h := 1; i+h < m && j+h < n && j-h >= 0; h++ {
				x := 0
				if j-h-1 >= 0 {
					x = s[i+h][j-h-1]
				}
				if s[i+h][j+h]-x == h*2+1 {
					ans++
				} else {
					break
				}
			}
		}
	}

	for i := m - 1; i > 0; i-- {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for h := 1; i-h >= 0 && j+h < n && j-h >= 0; h++ {
				x := 0
				if j-h-1 >= 0 {
					x = s[i-h][j-h-1]
				}
				if s[i-h][j+h]-x == h*2+1 {
					ans++
				} else {
					break
				}
			}
		}
	}

	return ans
}
