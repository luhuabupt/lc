package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hs struct{ sort.IntSlice }

func (h hs) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h hs) Swap(i, j int)       { h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i] }
func (h *hs) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hs) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type hi []int

func (h hi) Less(i, j int) bool  { return h[i] > h[j] }
func (h hi) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hi) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hi) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func main() {
	//p := [][]int{[]int{3, 10}, []int{1, 5}, []int{2, 6}}
	//fmt.Println(smallestChair(p, 0))
	fmt.Println(longestObstacleCourseAtEachPosition([]int{2, 2, 1}))
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))
	//fmt.Println(binarySearch([]int{2,4,6,8}, 9))
}

func binarySearch(arr []int, v int) int {
	l, r, pos := 0, len(arr)-1, 0
	for l <= r {
		mid := (l + r) / 2
		if v < arr[mid] {
			pos, r = mid, mid-1
		} else {
			pos, l = mid+1, mid+1
		}
	}
	return pos
}

// 优先队列
func smallestChair(times [][]int, targetFriend int) int {
	ans, arrive, leave, m := 0, [][]int{}, [][]int{}, map[int]int{}
	h := &hs{}
	for i := 0; i < len(times); i++ {
		heap.Push(h, i)
	}
	for k, arr := range times {
		arrive = append(arrive, []int{arr[0], k})
		leave = append(leave, []int{arr[1], k})
	}
	sort.Slice(arrive, func(i, j int) bool {
		return arrive[i][0] < arrive[j][0]
	})
	sort.Slice(leave, func(i, j int) bool {
		return leave[i][0] < leave[j][0]
	})
	i, j := 0, 0
	for {
		if arrive[i][0] < leave[j][0] {
			m[arrive[i][1]] = heap.Pop(h).(int)
			if arrive[i][1] == targetFriend {
				ans = m[arrive[i][1]]
				break
			}
			i++
		} else {
			heap.Push(h, m[leave[j][1]])
			j++
		}
	}
	return ans
}

func minStoneSum(piles []int, k int) int {
	h := &hs{}
	for _, v := range piles {
		heap.Push(h, v)
	}
	fmt.Println(h)
	for i := 0; i < k; i++ {
		x := heap.Pop(h).(int)
		fmt.Println(i, x)
		heap.Push(h, x-x/2)
		fmt.Println(h)
	}
	ans := 0
	for i := 0; i < len(piles); i++ {
		ans += heap.Pop(h).(int)
	}
	return ans
}

func minSwaps(s string) int {
	arr, n, idx, swap := []byte(s), len(s), 0, 0
	for i, j := 0, n-1; i < j; i++ {
		if arr[i] == '[' {
			idx++
		}
		if arr[i] == ']' {
			if idx == 0 {
				for arr[j] != '[' {
					j--
				}
				arr[i], arr[j] = '[', ']'
				idx++
				swap++
			} else {
				idx--
			}
		}
	}
	return swap
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	dp, ans := []int{}, []int{}
	for _, v := range obstacles {
		p := binarySearch(dp, v)
		ans = append(ans, p+1)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
	}
	return ans
}
