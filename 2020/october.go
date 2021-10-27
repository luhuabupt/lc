package main

import (
	"fmt"
	"sort"
)

func minimumDifference_(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	half := sum / 2

	n := len(nums) / 2
	a := make([][]int, n+1)
	b := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		a[i] = []int{}
		b[i] = []int{}
	}

	for i := 0; i < 1<<n; i++ {
		l, v, v2 := 0, 0, 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j]
				v2 += nums[n+j]
				l++
			}
		}
		a[l] = append(a[l], v)
		b[l] = append(b[l], v2)
	}

	for l := 0; l <= n; l++ {
		sort.Ints(a[l])
		sort.Ints(b[l])
	}

	ans := sum - 2*a[n][0]
	if ans < 0 {
		ans = -ans
	}
	//fmt.Println(a, b)

	for l := 0; l <= n; l++ {
		//fmt.Println(a[l], b[n-l])
		for i, j := 0, len(b[n-l])-1; i < len(a[l]); i++ {
			for a[l][i]+b[n-l][j] > half && j > 0 {
				//fmt.Println(i, j)
				ans = minAns(a[l][i]+b[n-l][j], sum, ans)
				j--
			}
			//fmt.Println(i, j)
			ans = minAns(a[l][i]+b[n-l][j], sum, ans)
		}
	}

	return ans
}

func minAns(a, sum, ans int) int {
	d := sum - 2*a
	if d < 0 {
		d = -d
	}
	if d < ans {
		return d
	}
	return ans
}

func minAbsDifference(nums []int, goal int) int {
	// [-2772,6927,4773,-2687,7167,-8995,2940,8869,526]
	//969621127 | 969590451 969589925
	n := len(nums) / 2
	a, b := []int{}, []int{}

	for i := 0; i < 1<<n; i++ {
		v := 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j]
			}
		}
		a = append(a, v)
	}
	for i := 0; i < 1<<(len(nums)-n); i++ {
		v := 0
		for j := 0; j < len(nums)-n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j+n]
			}
		}
		b = append(b, v)
	}

	sort.Ints(a)
	sort.Ints(b)

	//fmt.Println(a)
	//fmt.Println(b)

	ans := ab(a[0] - goal)

	for i, j := 0, len(b)-1; i < len(a); i++ {
		for j > 0 && a[i]+b[j] > goal {
			if ab(a[i]+b[j]-goal) < ans {
				ans = ab(a[i] + b[j] - goal)
			}
			j--
		}
		if ab(a[i]+b[j]-goal) < ans {
			ans = ab(a[i] + b[j] - goal)
		}
	}

	return ans
}

func ab(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 换位置保证nums1先遍历
	if len(nums2) < len(nums1) || len(nums2) == len(nums1) && nums2[len(nums2)-1] < nums1[len(nums1)-1] {
		return findMedianSortedArrays(nums2, nums1)
	}

	i, j, l, cur, prev := 0, 0, 1, 0, 0
	for i < len(nums1) && j < len(nums2) && l <= (len(nums1)+len(nums2)+2)/2 {
		if i < len(nums1) && nums1[i] <= nums2[j] {
			prev, cur = cur, nums1[i]
			i++
		} else {
			prev, cur = cur, nums2[j]
			j++
		}
		l++
	}

	for j < len(nums2) && l <= (len(nums1)+len(nums2)+2)/2 {
		prev, cur = cur, nums2[j]
		j++
		l++
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return (float64(prev) + float64(cur)) / 2
	} else {
		return float64(cur)
	}
}

func findJudge(n int, trust [][]int) int {
	m := map[int]map[int]bool{}
	t := map[int]bool{}
	for _, arr := range trust {
		if m[arr[1]] == nil {
			m[arr[1]] = map[int]bool{}
		}
		m[arr[1]][arr[0]] = true
		t[arr[0]] = true
	}
	for k, arr := range m {
		if len(arr) == n-1 && !arr[k] && !t[k] {
			return k
		}
	}
	return -1
}

type WordDictionary struct {
	ch    [26]*WordDictionary
	isEnd bool
}

func (this *WordDictionary) AddWord(word string) {
	x := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if x.ch[idx] == nil {
			x.ch[idx] = &WordDictionary{}
		}
		x = x.ch[idx]
	}
	x.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	x := this
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for _, v := range x.ch {
				if v == nil {
					continue
				}
				if i == len(word)-1 {
					if v.isEnd == true {
						return true
					}
				} else if v.Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			idx := word[i] - 'a'
			if x.ch[idx] == nil {
				return false
			}
			x = x.ch[idx]
		}
	}
	return x.isEnd
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	big := make([]int, n)
	ans := []int{}
	pos := map[int]int{}
	stack := []int{-1}

	for i := n - 1; i >= 0; i-- {
		v := nums2[i]
		pos[v], big[i] = i, -1

		stack = append(stack, v)
		x := len(stack) - 2
		for ; x >= 0 && v > stack[x]; x-- {
			stack[x] = v
			stack = stack[:x+1]
		}
		if x >= 0 {
			big[i] = stack[x]
		}
		//fmt.Println(i, stack, x)
	}

	for i := 0; i < len(nums1); i++ {
		ans = append(ans, big[pos[nums1[i]]])
	}

	return ans
}

func countConsistentStrings(allowed string, words []string) int {
	m := map[int32]bool{}
	for _, v := range allowed {
		m[v] = true
	}
	ans := 0
	for _, v := range words {
		flag := true
		for _, x := range v {
			if !m[x] {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	sum := make([]int, n)

	for i, v := range nums {
		if i == 0 {
			sum[i] = v
			continue
		}
		sum[i] = sum[i-1] + v
	}

	for i := 0; i < n; i++ {
		x := i * nums[i]
		if i > 0 {
			x -= sum[i-1]
		}

		x += sum[n-1] - sum[i]
		x -= (n - 1 - i) * nums[i]

		ans[i] = x
	}

	return ans
}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	sum := make([]int, n)
	sb := 0
	for i := 0; i < n; i++ {
		sum[i] = aliceValues[i] + bobValues[i]
		sb += bobValues[i]
	}
	sort.Ints(sum)
	a := 0
	for i, _ := range sum {
		if i&1 == 0 {
			a += sum[n-1-i]
		}
	}

	ans := 0
	if a > sb {
		ans = 1
	} else if a < sb {
		ans = -1
	}
	return ans
}

func minDeletions(s string) int {
	cnt := make([]int, 26)
	for _, v := range s {
		cnt[v-'a']++
	}
	sort.Ints(cnt)
	ans := 0
	for i := 24; i >= 0; i-- {
		for cnt[i] > 0 && cnt[i] >= cnt[i+1] {
			cnt[i]--
			ans++
		}
	}
	return ans
}

func maxProfit(inventory []int, orders int) int {
	mod := int(1e9) + 7
	n := len(inventory)

	sort.Ints(inventory)
	sum := 0
	for _, v := range inventory {
		sum += v
	}

	left := sum - orders
	avg := left / n
	extra := left % n

	i := 0
	for inventory[i] <= avg {
		for ; i < n && inventory[i] <= avg; i++ {
			extra += avg - inventory[i]
		}
		avg += extra / (n - i)
		extra %= n - i
	}

	ans := 0
	for ; i < n; i++ {
		if extra == n-i {
			avg++
		}
		ans += (inventory[i] - avg) * (inventory[i] + avg + 1) / 2
		ans %= mod
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{680754224, 895956841, 524658639, 3161416, 992716630, 7365440, 582714485, 422256708, 332815744, 269407767}, 707568720))
	//fmt.Println(maxProfit([]int{1,1,1,1000,10000}, 600))
}
