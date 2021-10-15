package main

func main() {

}

func construct2DArray(original []int, m int, n int) [][]int {
	ans := [][]int{}
	x := len(original)
	if m*n != x {
		return ans
	}
	k := 0
	for i := 0; i < m; i++ {
		one := []int{}
		for j := 0; j < n; j++ {
			one = append(one, original[k])
			k++
		}
		ans = append(ans, one)
	}
	return ans
}

func numOfPairs(nums []string, target string) int {
	n := len(nums)
	x := len(target)
	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if len(nums[i])+len(nums[j]) != x {
				continue
			}
			if nums[i]+nums[j] == target {
				ans++
			}
			if nums[j]+nums[i] == target {
				ans++
			}
		}
	}
	return ans
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return ctm(con(answerKey, k, 'T'), con(answerKey, k, 'F'))
}

func ctm(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func con(answerKey string, k int, x byte) int {
	n := len(answerKey)
	i, j := 0, 0
	for ; k > 0 && j < n; j++ {
		if answerKey[j] != x {
			if k > 0 {
				k--
			} else {
				break
			}
		}
	}
	ans := j
	j++
	for j < n {
		for j < n && answerKey[j] == x {
			j++
		}
		for answerKey[i] == x {
			i++
		}
		if j-i-1 > ans {
			ans = j - i - 1
		}
		i++
		j++
	}

	return ans
}

func waysToPartition(nums []int, k int) int {
	n := len(nums)
	l, r := map[int]int{}, map[int]int{}

	tot, sum, ans := 0, 0, 0
	for _, v := range nums {
		tot += v
	}
	for i, v := range nums {
		if i == n-1 {
			continue
		}
		sum += v
		if sum*2 == tot {
			ans++
		}
	}

	sum = 0
	for i := n - 1; i >= 0; i-- {
		sum += nums[i]
		r[sum*2]++
	}

	sum = 0
	for i := 0; i < n; i++ { // 改i
		d := k - nums[i]
		r[(tot-sum)*2]--

		if l[tot+d]+r[tot+d] > ans {
			ans = l[tot+d] + r[tot+d]
		}

		sum += nums[i]
		l[sum*2]++
	}

	return ans
}
