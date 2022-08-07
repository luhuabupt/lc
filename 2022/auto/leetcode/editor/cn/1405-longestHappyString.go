package main

import "fmt"

func main() {
	fmt.Println(longestDiverseString(0, 4, 7))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestDiverseString(a int, b int, c int) string {
	arr := []int{a, b, c}
	ans := []byte{}

	max := func(a []int, ex int) int {
		mx := -1
		for i, v := range arr {
			if ex == i {
				continue
			}
			if mx == -1 || v > arr[mx] {
				mx = i
			}
		}
		return mx
	}

	for {
		ex := -1
		if len(ans) > 1 {
			if ans[len(ans)-1] == ans[len(ans)-2] {
				ex = int(ans[len(ans)-1] - 'a')
			}
		}
		x := max(arr, ex)
		if arr[x] == 0 {
			break
		}
		ans = append(ans, byte(x+'a'))
		arr[x]--
	}

	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
