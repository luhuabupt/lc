package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(1234)
}

func executeInstructions(n int, startPos []int, s string) []int {
	dr := map[byte][]int{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}
	ans := []int{}
	for i := 0; i < len(s); i++ {
		x, y, d, ii := startPos[0]+dr[s[i]][0], startPos[1]+dr[s[i]][1], 0, i
		for ; x >= 0 && x < n && y >= 0 && y < n && ii < len(s); ii++ {
			d++
			x += dr[s[ii]][0]
			y += dr[s[ii]][1]
		}
		ans = append(ans, d)
	}
	return ans
}

func getDistances(arr []int) []int64 {
	cnt := map[int][][]int{} // 0:i 1:sum
	idx := map[int]int{}
	for i, v := range arr {
		if _, ok := cnt[v]; !ok {
			cnt[v] = [][]int{{i, i}}
		} else {
			cnt[v] = append(cnt[v], []int{i, i + cnt[v][len(cnt[v])-1][1]})
		}
		idx[v] = 0
	}

	ans := []int64{}
	for i, v := range arr {
		val := 0
		if len(cnt[v]) > 0 {
			ci, pi := idx[v], idx[v]-1
			sum := cnt[v]
			if ci > 0 {
				val += ci*i - sum[pi][1]
			}
			if ci < len(sum)-1 {
				val += sum[len(sum)-1][1]
				val -= sum[ci][1]
				val -= i * (len(sum) - 1 - ci)
			}
			idx[v]++
		}
		ans = append(ans, int64(val))
	}
	return ans
}

func recoverArray(nums []int) []int {
	n := len(nums)
	if n == 2 {
		return []int{(nums[0] + nums[1]) / 2}
	}

	sort.Ints(nums)
	k := []int{}
	vis := map[int]bool{}
	for i := 0; i <= n/2; i++ {
		if nums[i]-nums[0] > 0 && !vis[nums[i]-nums[0]] {
			k = append(k, (nums[i]-nums[0])/2)
			vis[nums[i]-nums[0]] = true
		}
	}

	for _, kk := range k {
		ans := []int{}
		cnt := map[int]int{}
		for _, v := range nums {
			cnt[v]++
		}
		for _, x := range nums {
			if cnt[x] > 0 {
				if cnt[x+2*kk] == 0 {
					break
				}
				ans = append(ans, x+kk)
				cnt[x]--
				cnt[x+2*kk]--
			}
		}
		if len(ans) == n/2 {
			return ans
		}
	}
	return []int{}
}
