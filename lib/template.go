package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func int2Arr(a int) []int {
	ans := []int{}
	for x := a; x >= 0; x /= 10 {
		ans = append(ans, a)
	}
	for i := 0; i < len(ans); i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}
