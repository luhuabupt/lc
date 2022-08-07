package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func rotatedDigits(n int) (ans int) {
	get := func(v int) (ans []int) {
		for v > 0 {
			ans = append(ans, v%10)
			v /= 10
		}
		return
	}

	do := map[int]int{2: 2, 5: 2, 6: 2, 9: 2, 0: 1, 1: 1, 8: 1}
	for i := 1; i <= n; i++ {
		if do[i%10] == 0 {
			continue
		}

		re := get(i)
		flag := false
		for _, x := range re {
			if do[x] == 0 {
				flag = false
				break
			}
			if do[x] == 2 {
				flag = true
			}
		}
		if flag {
			ans++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
