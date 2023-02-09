package main

import (
	"fmt"
	"sync"
)

func main() {
	slice()
}

func slice() {
	a := []int{1, 2, 3, 4}
	b := a[1:]
	b[1] = 10
	fmt.Println(a, b)

	b = append(b, 5)
	fmt.Println(a, b)
	b[1] = 20
	fmt.Println(a, b)

	c := make([]int, 3, 5)
	d := c[1:3:5]
	c[1] = 20
	d = append(d, 20)
	//d = append(d, 20)
	fmt.Println(c, d)

	d[1] = 30
	fmt.Println(c, d)
	c = append(c, 50)
	fmt.Println(c, d)

}

// 1 ~ 100
func sum(n int) int {
	ans := 0
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	add := func(i int) {
		lock.Lock()
		ans += i
		wg.Done()
		lock.Unlock()
	}

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go add(i)
	}

	wg.Wait()
	return ans
}

// 给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
//
//在 S 上反复执行重复项删除操作，直到无法继续删除。
//
//在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。
//
//示例：
//
//输入："abbaca"
//输出："ca"
//解释：
//例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
//
//提示：
//
//1 <= S.length <= 20000
//S 仅由小写英文字母组成。

func delDuplicate(s string) string {
	stack := []byte{}
	for _, v := range s {
		if len(stack) > 0 && stack[len(stack)-1] == byte(v) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(v))
		}
	}
	return string(stack)
}

func atoi(str string) (int, string) {
	if str == "" {
		return 0, "empty error"
	}
	if len(str) > 21 {
		return 0, "string len too long"
	}

	// 负数
	neg := false
	if str[0] == '-' {
		neg = true
		str = str[1:]
	}

	// 数字校验
	for _, v := range str {
		if v < '0' || v > '9' {
			return 0, "not number"
		}
	}

	// 特判 -1<<63
	if str == "9223372036854775808" {
		return -1 << 63, ""
	}

	// 范围：-(1<<63 - 1)~ 1<<63-1
	max := 1<<63 - 1

	ans := 0
	for _, v := range str {
		if float64(ans) > float64(max/10) {
			return 0, "over limit"
		}
		ans *= 10
		if ans > max-(int(v-'0')) {
			return 0, "over limit"
		}
		ans += int(v - '0')
	}

	if neg {
		return -ans, ""
	}
	return ans, ""
}
