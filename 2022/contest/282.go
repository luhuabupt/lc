package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(minimumFinishTime([][]int{{2,3}, {3,4}}, 5, 4))
	//fmt.Println(minimumFinishTime([][]int{{1, 10}, {2, 2}, {3, 4}}, 6, 5))
	//fmt.Println(minimumFinishTime([][]int{{1, 2}}, 1,1))

	fmt.Println(minimumTime([]int{1e7}, int(1e7)))

	fmt.Println(1<<32 - 1)
	var a int
	a = 1 << 61
	fmt.Println(a)
	fmt.Println(math.MaxInt64)
	fmt.Println(int(1e14))
}

func minimumTime(time []int, totalTrips int) int64 {
	return int64(sort.Search(totalTrips*1e7, func(x int) bool {
		tot := 0
		for _, t := range time {
			tot += x / t
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}

func minimumFinishTime(t [][]int, ct int, lp int) int {
	do := make([][]int, lp)
	for _, v := range t {
		x, s := v[0], v[0]+ct
		do[0] = append(do[0], s)
		for i := 1; i < lp; i++ {
			x *= v[1]
			s += x
			do[i] = append(do[i], s)
			if x >= v[0]+ct {
				break
			}
		}
	}

	//fmt.Println(do)

	dp := make([]int, lp)
	for i := 0; i < lp; i++ {
		x := 1<<31 - 1
		if len(do[i]) > 0 {
			for _, v := range do[i] {
				if v < x {
					x = v
				}
			}
		}
		for j := 0; j < i; j++ {
			if dp[j]+dp[i-j-1] < x {
				x = dp[j] + dp[i-j-1]
			}
		}
		dp[i] = x
	}

	//fmt.Println(dp)

	return dp[lp-1] - ct
}
