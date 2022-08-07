package main

import "fmt"

func main() {
	fmt.Println(monotoneIncreasingDigits(1234))
	fmt.Println(monotoneIncreasingDigits(332))
}

//leetcode submit region begin(Prohibit modification and deletion)
func monotoneIncreasingDigits(n int) int {
	arr := []int{}
	for n > 0 {
		arr = append([]int{n % 10}, arr...)
		n /= 10
	}

	pre := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			pre = i
		} else if arr[i] == arr[i-1] {

		} else if arr[i] < arr[i-1] {
			arr[pre] = arr[pre] - 1
			for x := pre + 1; x < len(arr); x++ {
				arr[x] = 9
			}
		}
	}

	ans := 0
	for i, v := range arr {
		if i == 0 && v == 0 {
			continue
		}
		ans *= 10
		ans += v
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
