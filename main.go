package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 4}
	a = []int{6, 8, 1, 8, 6, 5, 6, 11, 17}
	fmt.Println(numberOfGoodSubsets(a))
	//fmt.Println(numberOfGoodSubsets_(a))
}

func numberOfGoodSubsets(nums []int) int {
	no := map[int]int{0: 0, 1: 0, 2: 1, 3: 2, 4: 0, 5: 4, 6: 3, 7: 8, 8: 0, 9: 0,
		10: 5, 11: 16, 12: 0, 13: 32, 14: 9, 15: 6, 16: 0, 17: 64, 18: 0, 19: 128,
		20: 0, 21: 10, 22: 17, 23: 256, 24: 0, 25: 0, 26: 33, 27: 0, 28: 0, 29: 512, 30: 7}
	c, dp, one, mod, ans := map[int]int{}, map[int]int{}, 1, int(1e9+7), 0
	for _, v := range nums {
		if v == 1 {
			one = (one << 1) % mod
			continue
		}
		if no[v] == 0 {
			continue
		}
		x := no[v]
		c[x]++
	}
	fmt.Println(c)

	for i := 1; i < 1024; i++ {
		for j := 1; j < i; j++ {
			if i&j == 0 && c[j] > 0 {
				if dp[i] == 0 {
					dp[i] = 1
				}
				dp[i] *= c[j]
				dp[i] %= mod
				//dp[i|j] += dp[i]*dp[j]
				//dp[i|j] %= mod
			}
		}
	}
	for k, v := range dp {
		if v == 0 {
			delete(dp, k)
		}
		ans += v
		ans %= mod
	}
	fmt.Println(dp)
	ans = ans * one
	ans %= mod
	return ans
}

func numberOfGoodSubsets_(nums []int) int {
	p := 1000000007
	ps := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	ss := make([]int, 1024)
	for _, x := range nums {
		u, e := 0, 1
		for _, pi := range ps {
			if x%pi == 0 {
				u |= e
				x /= pi
			}
			if x%pi == 0 {
				u = 0
				break
			}
			e <<= 1
		}
		if u != 0 {
			ss[u]++
		}
	}
	dp := make([]int, 1024)
	dp[0] = 1
	for i := 0; i < 1024; i++ {
		if ss[i] > 0 {
			for j := 0; j < 1024; j++ {
				if i&j == 0 {
					dp[i|j] = (dp[i|j] + ss[i]*dp[j]%p) % p
				}
			}
		}
	}
	c := 1
	for _, pi := range nums {
		if pi == 1 {
			c = (c << 1) % p
		}
	}
	ans := 0
	for i := 1; i < 1024; i++ {
		ans = (ans + dp[i]) % p
	}
	return (ans * c) % p
}
