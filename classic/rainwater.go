package classic

// https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	n := len(height)
	l, r := make([]int, n), make([]int, n) // 不包含i的最大值

	l[0], r[n-1] = 0, 0
	for i := 1; i < n; i++ {
		l[i] = l[i-1]
		if height[i-1] > l[i-1] {
			l[i] = height[i-1]
		}
	}

	for i := n - 2; i >= 0; i--{
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
func trap_(height []int) int {
	n := len(height)

	l, r := height[0], height[n-1] // 不包含i, j
	i, j, ans := 1, n-2, 0
	for i < j {
		if l < r {
			i++
			if height[i] < l {
				ans += l - height[i]
			} else {
				l = height[i]
			}
		} else {
			j--
			if height[j] < r {
				ans += r - height[i]
			} else {
				ans = height[j]
			}
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
