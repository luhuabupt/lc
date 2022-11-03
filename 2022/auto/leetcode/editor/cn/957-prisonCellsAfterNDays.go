package main

import "fmt"

func main() {
	//Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
	//Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
	//Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
	//Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
	//Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
	//Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
	//Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
	//Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}

//leetcode submit region begin(Prohibit modification and deletion)
func prisonAfterNDays(a []int, n int) []int {
	d := make([]int, 1<<9)
	dt := make([]int, 1<<9)
	for i, _ := range d {
		d[i] = -1
	}

	start := 0
	loop := 0
	var dfs func(i, t int)
	dfs = func(i, t int) {
		//fmt.Printf("%d %b \n", t, i)
		if d[i] > 0 {
			start = d[i]
			loop = t - d[i]
			return
		}
		d[i] = t
		dt[t] = i
		nt := 0
		for j := 1; j <= 6; j++ {
			x := (i >> (j - 1) & 1) ^ ((i >> (j + 1)) & 1) ^ 1
			nt += x << j
			//fmt.Println(" ", j, x, x<<j, nt)
		}
		dfs(nt, t+1)
	}

	s := 0
	for i, v := range a {
		s += v * (1 << i)
	}
	dfs(s, 0)

	final := 0
	if n < start {
		final = dt[n]
	} else {
		n -= start
		n %= loop
		final = dt[start+n]
	}

	ans := []int{}
	for i := 0; i < 8; i++ {
		ans = append(ans, (final>>i)&1)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
