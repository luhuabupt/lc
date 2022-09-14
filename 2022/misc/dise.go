package main

import "fmt"

func main() {
	a := 6
	b := 24
	fmt.Println(fmt.Sprintf("%d, %d, %f ", a, b, cal(a, b)))
}

func cal(a, b int) float64 {
	max := a * 6
	d := make([][]int, a+1) // 第i个, 和为v
	d[0] = make([]int, max+1)
	d[0][0] = 1

	for i := 1; i <= a; i++ {
		d[i] = make([]int, max+1)

		for v := 1; v <= 6; v++ {
			for j := max; j >= i; j-- {
				if j-v >= 0 {
					d[i][j] += d[i-1][j-v]
				}
			}
		}
	}

	for i := 1; i <= max; i++ {
		//fmt.Println(i, d[a][i])
	}

	sum := 0
	for i := 1; i <= max; i++ {
		sum += d[a][i]
	}
	//fmt.Println("sum:", sum)

	ans := 0
	for i := b; i <= max; i++ {
		ans += d[a][i]
	}

	return float64(ans) / float64(sum)
}
