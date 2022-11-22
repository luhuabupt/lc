package main

import "fmt"

// count-ways-to-make-array-with-product 生成乘积数组的方案数  2022-10-23 13:08:30
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
func waysToFillArray(q [][]int) []int {
	a := []int{}
	d := make([][]int, 1e4+1)
	//M := int(1e9 + 7)
	ans := make([]int, len(q))

loop:
	for i := 2; i < 101; i++ {
		for j := 2; j < i && j < 11; j++ {
			if i%j == 0 {
				continue loop
			}
		}
		a = append(a, i)
	}

	for i, v := range q {
		_, k := v[0], v[1]
		if k == 1 {
			ans[i] = 1
			continue
		}

		t := k

	l2:
		for t > 1 {
			for _, x := range a {
				if t%x == 0 {
					t /= x
					d[k] = append(d[k], x)
					continue l2
				}
			}
			break
		}

		if t > 1 {
			d[k] = append(d[k], t)
		}

		//dp := make([]int, n)

	}

	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
