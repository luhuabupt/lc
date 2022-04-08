package main

import (
	"fmt"
)

func main() {
	//fmt.Println("")
	//fmt.Println(sumScores("babab"))
	//fmt.Println(sumScores("azbazbzaz"))
	fmt.Println(sumScores("")) // 40765
}

func minBitFlips(start int, goal int) (ans int) {
	for i := 0; i < 32; i++ {
		x := 1 << i
		if x&start != x&goal {
			ans++
		}
	}
	return ans
}

func triangularSum(a []int) int {
	n := len(a)
	if n == 1 {
		return a[0]
	}
	nn := make([]int, n-1)
	for i := 1; i < n; i++ {
		nn[i-1] = (a[i-1] + a[i]) % 10
	}
	return triangularSum(nn)
}

func numberOfWays(s string) int64 {
	n := len(s)
	sa, sb := 0, 0
	a, b := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			sa++
		} else {
			sb++
		}
	}

	ans := int64(0)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			a++
			ans += int64(b) * int64(sb-b)
		} else {
			b++
			ans += int64(a) * int64(sa-a)
		}
	}

	return ans
}

func sumScores(ss string) int64 {
	do := 0
	calcMaxMatchLengths := func(s []byte) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				do++
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}

	kmpSearch := func(text, pattern []byte) (pos int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for _, v := range text {
			for c > 0 && pattern[c] != v {
				do++
				c = match[c-1]
			}
			if pattern[c] == v {
				c++
			}
			if c == lenP {
				pos++
				c = match[c-1] // 不允许重叠时 c = 0
			}
		}
		return
	}

	s := []byte(ss)
	n := len(s)
	ans := 0
	pre := 0
	for i := n - 1; i >= 0; i-- {
		x := 1
		if i < n-1 {
			x = kmpSearch(s, s[:i+1])
		}
		//fmt.Println(i, s[:i+1], x,  (i+1) * (x - pre), x, pre, x-pre)
		ans += (i + 1) * (x - pre)
		pre = x
	}

	fmt.Println(n, do)
	return int64(ans)
}

func generateNext(p string) []int {
	m := len(p)
	next := make([]int, m, m)
	next[0] = -1
	next[1] = 0
	i, j := 0, 1 // 前缀子串、后缀子串起始位置
	// 因为是通过最长可匹配前缀子串计算，所以 j 的值最大不会超过 m-1
	for j < m-1 {
		if i == -1 || p[i] == p[j] {
			i++
			j++
			// 设置当前最长可匹配前缀子串结尾字符下标
			next[j] = i
		} else {
			// 如果 p[i] != p[j]，回到上一个最长可匹配前缀子串结尾字符下标
			i = next[i]
		}
	}
	return next
}

// KMP 算法实现函数
func kmpSearch(s, p string) int {
	n := len(s)             // 主串长度
	m := len(p)             // 模式串长度
	next := generateNext(p) // 生成 next 数组
	i, j := 0, 0
	for i < n && j < m {
		// 如果主串字符和模式串字符不相等，
		// 更新模式串坏字符下标位置为好前缀最长可匹配前缀子串尾字符下标
		// 然后从这个位置重新开始与主串匹配
		// 相当于前面提到的把模式串向后移动 j - k 位
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		// 完全匹配，返回下标位置
		return i - j
	} else {
		return -1
	}
}

// 基于 KMP 算法实现查找字符串子串函数
func strStrV2(haystack, needle string) int {
	// 子串长度=0
	if len(needle) == 0 {
		return 0
	}
	//主串长度=0，或者主串长度小于子串长度
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	// 子串长度=1时单独判断
	if len(needle) == 1 {
		i := 0
		for ; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				return i
			}
		}
		return -1
	}
	// 其他情况走 KMP 算法
	return kmpSearch(haystack, needle)
}

//func getNext(s string) []int {
//	next:=make([]int,len(s))
//	next[0]=-1
//	i,j:=0,-1
//	for ; i< len(s)-1;  {
//		if j==-1||s[i]==s[j] {
//			i++
//			j++
//			next[i]=j
//		}else {
//			j=next[j]
//		}
//	}
//	return next
//}
//func KMPSearch(target string,text string) int {
//	i,j:=0,0
//	next := getNext(target)
//	for ; j< len(text);  {
//		if i==len(target)-1&&target[i]==text[j] {
//			return j-i
//		}
//		if j==-1||target[i]==text[j] {
//			i++
//			j++
//		}else {
//			i=next[i]
//		}
//	}
//	return -1
//}
