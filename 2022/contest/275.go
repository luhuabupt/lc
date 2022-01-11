package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(wordCount([]string{"ab", "a"}, []string{"abc", "abcd"}))
	fmt.Println(wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
}

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	for _, arr := range matrix {
		cnt := make([]int, n)
		for _, v := range arr {
			if v > n {
				return false
			}
			cnt[v-1]++
		}
		for i := 1; i <= n; i++ {
			if cnt[i-1] != 1 {
				return false
			}
		}
	}

	for j := 0; j < n; j++ {
		cnt := make([]int, n)
		for i := 0; i < n; i++ {
			if matrix[i][j] > n {
				return false
			}
			cnt[matrix[i][j]-1]++
		}
		for i := 1; i <= n; i++ {
			if cnt[i-1] != 1 {
				return false
			}
		}
	}
	return true
}

func minSwaps(nums []int) int {
	n := len(nums)
	cnt, max := 0, 0
	for _, v := range nums {
		cnt += v
	}

	sum := []int{0}
	for i := 0; i < n; i++ {
		sum = append(sum, nums[i]+sum[i])
	}
	for i := 0; i < cnt; i++ {
		sum = append(sum, sum[len(sum)-1]+nums[i])
	}

	for i := cnt; i < len(sum); i++ {
		if sum[i]-sum[i-cnt] > max {
			max = sum[i] - sum[i-cnt]
		}
	}
	return cnt - max
}

func wordCount(startWords []string, targetWords []string) int {
	get := func(a string) string {
		ans := []int{}
		for _, v := range a {
			ans = append(ans, int(v))
		}
		sort.Ints(ans)
		ba := []byte{}
		for _, v := range ans {
			ba = append(ba, byte(v))
		}
		return string(ba)
	}

	m := map[string]bool{}
	for _, v := range startWords {
		m[get(v)] = true
	}

	ans := 0
	for _, s := range targetWords {
		s = get(s)
		for i := 0; i < len(s); i++ {
			x := []byte{}
			for j := 0; j < len(s); j++ {
				if i != j {
					x = append(x, s[j])
				}
			}
			if m[string(x)] {
				ans++
				break
			}
		}
	}

	return ans
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	type pair struct {
		p, g int
	}
	arr := []*pair{}
	for i, v := range plantTime {
		arr = append(arr, &pair{v, growTime[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].g > arr[j].g
	})

	use, max := 0, 0
	for _, v := range arr {
		use += v.p
		if use+v.g > max {
			max = use + v.g
		}
	}
	return max
}
