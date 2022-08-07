package main

import "strconv"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func evalRPN(a []string) int {
	st := []int{}
	for _, v := range a {
		n := len(st)
		if v == "+" {
			st[n-2] = st[n-1] + st[n-2]
			st = st[:n-1]
		} else if v == "-" {
			st[n-2] = st[n-2] - st[n-1]
			st = st[:n-1]
		} else if v == "*" {
			st[n-2] = st[n-1] * st[n-2]
			st = st[:n-1]
		} else if v == "/" {
			st[n-2] = st[n-2] / st[n-1]
			st = st[:n-1]
		} else {
			x, _ := strconv.Atoi(v)
			st = append(st, x)
		}
	}

	return st[0]
}

//leetcode submit region end(Prohibit modification and deletion)
