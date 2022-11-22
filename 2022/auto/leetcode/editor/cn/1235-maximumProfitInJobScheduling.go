package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// maximum-profit-in-job-scheduling 规划兼职工作  2022-10-22 10:09:59
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
type ht [][2]int

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i][0] < h[j][0] } // 小顶堆
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	arr := make([][3]int, len(startTime))
	for i, v := range startTime {
		arr[i] = [3]int{v, endTime[i], profit[i]}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	arr = append(arr, [3]int{2e9, 2e9 + 1, 0})

	h := &ht{}
	cur := 0
	for _, v := range arr {
		for h.Len() > 0 {
			x := heap.Pop(h).([2]int)
			if x[0] > v[0] {
				heap.Push(h, x)
				break
			}
			if x[1] > cur {
				cur = x[1]
			}
		}
		heap.Push(h, [2]int{v[1], cur + v[2]}) // e p
	}

	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
