package classic

// https://leetcode-cn.com/problems/trapping-rain-water/
func trap_(height []int) int { // 双单调栈
	n := len(height)
	l, r := make([]int, n), make([]int, n) // 不包含i的最大值

	l[0], r[n-1] = 0, 0
	for i := 1; i < n; i++ {
		l[i] = l[i-1]
		if height[i-1] > l[i-1] {
			l[i] = height[i-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1]
		if height[i+1] > r[i+1] {
			r[i] = height[i+1]
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		x := twoMin(l[i], r[i])
		if x > height[i] {
			ans += x - height[i]
		}
	}

	return ans
}

// 双指针
func trap(height []int) int {
	n := len(height)
	l, r := 0, 0 // 包含i, j
	i, j, ans := 0, n-1, 0

	for i < j {
		if height[i] > l {
			l = height[i]
		}
		if height[j] > r {
			r = height[j]
		}
		if l < r {
			ans += l - height[i]
			i++
		} else {
			ans += r - height[j]
			j--
		}
	}

	return ans
}

func twoMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
