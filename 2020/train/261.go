package main

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	m := len(rolls)
	for _, v := range rolls {
		sum += v
	}

	tot := (m+n)*mean - sum
	if tot < n || tot > 6*n {
		return []int{}
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if n-i-1 == tot-1 {
			ans = append(ans, 1)
			tot--
		} else if n-i-1 <= tot-6 {
			ans = append(ans, 6)
			tot -= 6
		} else {
			ans = append(ans, tot-n+1+i)
			tot -= tot - n + 1 + i
		}
	}
	return ans
}

func stoneGameIX(stones []int) bool {
	ct := [3]int{}
	for _, v := range stones {
		ct[v%3]++
	}
	if ct[1] == 0 || ct[2] == 0 {
		return false
	}

	return check(ct[0], ct[1], ct[2], len(stones)) || check(ct[0], ct[2], ct[1], len(stones))
}

func check(l, o, t, n int) bool {
	// o 开头
	l += 1 + 2*min(o-1, t)
	if o-1 > t {
		l++
	}

	return l%2 == 1 && l < n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	n := len(s)
	stack := []byte{s[0]}

	letCnt, cnt := make([]int, n), 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == letter {
			cnt++
		}
		letCnt[i] = cnt
	}

	cnt = 0
	for i := 1; i < n; i++ {
		if n-1-i == k-len(stack) {
			stack = append(stack, s[i])
			continue
		}
		if s[i] == letter { // 必须补letter
			if cnt+letCnt[i] == repetition {
				for x := 0; x < letCnt[i]; x++ {
					stack = append(stack, letter)
					return string(stack)
				}
			}
		}
		if s[i] < stack[len(stack)-1] {
			for top := len(stack) - 1; top >= 0 && top >= k+i-n && s[i] < stack[top]; top-- {
				if stack[top] == letter {
					if cnt+letCnt[i] == repetition { // letter不够了 不能出letter
						break
					}
					cnt--
				}
				stack[top] = s[i]
				stack = stack[:top+1]
			}
			if s[i] == letter {
				cnt++
			}
		} else if len(stack) < k {
			stack = append(stack, s[i])
			if s[i] == letter {
				cnt++
			}
		}
	}

	return string(stack)
}

func smallestSubsequence_(s string, k int, letter byte, repetition int) string {
	n := len(s)
	m := make([]map[byte]int, n)
	m[0] = make(map[byte]int)
	m[0][s[0]] = 1
	for i := 1; i < n; i++ {
		m[i] = make(map[byte]int)
		for x := byte('a'); x < 'z'; x++ {
			if m[i-1][x] > 0 {
				m[i][x] = m[i-1][x]
			}
		}
		m[i][s[i]]++
	}
	letCnt := m[n-1][letter]

	pos := map[byte][]int{}
	for i := byte('a'); i < 'z'; i++ {
		pos[i] = []int{0}
	}
	for i := 0; i < n; i++ {
		pos[s[i]] = append(pos[s[i]], i)
	}

	fmt.Println(m)
	fmt.Println(pos)

	ans := []byte{}

	l, r := 0, pos[letter][letCnt-repetition+1]

	for k > 0 {
		fmt.Println(l, r)
		if k == repetition {
			for x := repetition; x > 0; x-- {
				ans = append(ans, letter)
			}
			return string(ans)
		}
		if k == n-l {
			for x := l; x < n; x++ {
				ans = append(ans, s[x])
			}
			return string(ans)
		}
		for i := byte('a'); i <= 'z'; i++ {
			if m[r][i]-m[l][i] > 0 {
				if i == letter {
					repetition--
				}
				ans = append(ans, i)
				l, r = pos[i][m[l][i]+1]+1, pos[letter][letCnt-k+1]
				break
			}
		}
		k--
	}

	return string(ans)
}
