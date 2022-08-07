package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(countSubarrays([]int{1, 1, 1}, int64(5)))
	//fmt.Println(countSubarrays([]int{2, 1, 4, 3, 5}, int64(10)))
	fmt.Println(matchReplacement("llllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrllllllllllllllllllllllllllllllllrlllllllllllllllllllllllllllllllll", "lllllllllllllllllllllllllllllllll", [][]byte{{'l', 'a'}, {'l', 'b'}, {'l', 'c'}, {'l', 'd'}, {'l', 'e'}, {'l', 'f'}, {'l', 'g'}, {'l', 'h'}, {'l', 'i'}, {'l', 'j'}, {'l', 'k'}, {'l', 'm'}, {'l', 'n'}, {'l', 'o'}, {'l', 'p'}, {'l', 'q'}, {'l', 's'}, {'l', 't'}, {'l', 'u'}, {'l', 'v'}, {'l', 'w'}, {'l', 'x'}, {'l', 'y'}, {'l', 'z'}, {'l', '0'}, {'l', '1'}, {'l', '2'}, {'l', '3'}, {'l', '4'}, {'l', '5'}, {'l', '6'}, {'l', '7'}, {'l', '8'}, {'l', '9'}, {'r', 'a'}, {'r', 'b'}, {'r', 'c'}, {'r', 'd'}, {'r', 'e'}, {'r', 'f'}, {'r', 'g'}, {'r', 'h'}, {'r', 'i'}, {'r', 'j'}, {'r', 'k'}, {'r', 'm'}, {'r', 'n'}, {'r', 'o'}, {'r', 'p'}, {'r', 'q'}, {'r', 's'}, {'r', 't'}, {'r', 'u'}, {'r', 'v'}, {'r', 'w'}, {'r', 'x'}, {'r', 'y'}, {'r', 'z'}, {'r', '0'}, {'r', '1'}, {'r', '2'}, {'r', '3'}, {'r', '4'}, {'r', '5'}, {'r', '6'}, {'r', '7'}, {'r', '8'}, {'r', '9'}}))
}

func matchReplacement(s string, sub string, ma [][]byte) bool {
	do := map[byte]map[byte]bool{}
	for _, v := range ma {
		if do[v[0]] == nil {
			do[v[0]] = map[byte]bool{}
		}
		do[v[0]][v[1]] = true
	}
	// fmt.Println(do)
	n, m := len(sub), len(s)

loop:
	for i := 0; i <= m-n; i++ {
		for j := 0; j < n; j++ {
			if s[i+j] != sub[j] {
				//fmt.Println(i, j, string(s[i+j]), string(sub[j]), do[sub[j]])
				if do[sub[j]] == nil || !do[sub[j]][s[i+j]] {
					continue loop
				}
			}
		}
		return true
	}
	return false
}

func successfulPairs(s []int, p []int, suc int64) []int {
	sort.Ints(p)
	ans := make([]int, len(s))
	for i, v := range s {
		ans[i] = len(p) - sort.Search(len(p), func(j int) bool {
			return int64(v*p[j]) >= suc
		})
	}
	return ans
}

func countSubarrays(a []int, k int64) int64 {
	ans := 0
	n := len(a)
	dp := make([]int, n)
	s := 0
	for i, v := range a {
		s += v
		dp[i] = s
	}

	for i, _ := range a {
		x := sort.Search(i+1, func(j int) bool {
			sp := dp[i]
			if j > 0 {
				sp -= dp[j-1]
			}
			if int64((i-j+1)*sp) < k {
				return true
			}
			return false
		})
		fmt.Println(i, x)
		ans += (i - x + 1)
	}
	return int64(ans)
}

func strongPasswordCheckerII(s string) bool {
	if len(s) < 8 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return false
		}
	}
	a, b, c, d := false, false, false, false
	sb := "!@#$%^&*()-+"
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			a = true
		}
		if v >= 'A' && v <= 'Z' {
			b = true
		}
		if v >= '0' && v <= '9' {
			c = true
		}
		for _, x := range sb {
			if v == x {
				d = true
			}
		}
	}
	return a && b && c && d
}

type TextEditor struct {
	a   []byte
	idx int
}

func Constructor() TextEditor {
	return TextEditor{[]byte{}, -1}
}

func (t *TextEditor) AddText(text string) {
	to := t.a[t.idx+1:]
	t.a = t.a[:max(t.idx, 0)]
	t.a = append(t.a, text...)
	t.idx = len(t.a) - 1
	t.a = append(t.a, to...)
}

func (t *TextEditor) DeleteText(k int) int {
	ans := min(k, t.idx+1)
	to := t.a[t.idx+1:]
	t.a = t.a[:max(t.idx-k, 0)]
	t.idx = len(t.a) - 1
	t.a = append(t.a, to...)
	return ans
}

func (t *TextEditor) CursorLeft(k int) string {
	t.idx = max(0, t.idx-k)
	ans := t.a[max(t.idx-10, 0):t.idx]
	return string(ans)
}

func (t *TextEditor) CursorRight(k int) string {
	t.idx = max(len(t.a)-1, t.idx+k)
	ans := t.a[max(t.idx-10, 0):t.idx]
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
