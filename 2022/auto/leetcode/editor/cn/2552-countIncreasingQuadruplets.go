package main

import "fmt"

func main() {
	fmt.Println(countQuadruplets([]int{1, 3, 2, 4, 5}))
	fmt.Println(countQuadruplets([]int{1, 2, 3, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countQuadruplets(a []int) int64 {
	n := len(a)

	less := make([][]int, n)
	for i, _ := range a {
		if i == 0 {
			less[i] = make([]int, n+1)
			continue
		}
		less[i] = append([]int{}, less[i-1]...)
		for j := a[i-1] + 1; j <= n; j++ {
			less[i][j]++
		}
	}

	great := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if a[j] > a[i] {
				great[i]++
			}
		}
	}

	//fmt.Println(less)
	//fmt.Println(great)

	ans := 0
	for i := 1; i < n-2; i++ {
		t := 0
		for j := i + 1; j < n-1; j++ {
			if a[j] > a[i] {
				t++
			}
			if a[j] < a[i] {
				//fmt.Println(i, j, less[i][a[j]], great[i], t)
				ans += less[i][a[j]] * (great[i] - t)
			}
		}
	}

	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
