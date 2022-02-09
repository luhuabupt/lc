package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minimumDeviation([]int{3, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)

type p struct {
	i, v int
}
type h []p

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i].v < h[j].v } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.(p)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func minimumDeviation(nums []int) int {
	n := len(nums)

	arr := make([][]int, n)
	for i, v := range nums {
		if v&1 == 1 {
			arr[i] = []int{v + v, v}
		} else {
			for x := v; ; x >>= 1 {
				arr[i] = append(arr[i], x)
				if x&1 == 1 {
					break
				}
			}
		}
	}
	pop := func(i int) int {
		if len(arr[i]) == 0 {
			return 0
		}
		x := arr[i][len(arr[i])-1]
		arr[i] = arr[i][:len(arr[i])-1]
		return x
	}

	h, max, ans := &h{}, -1<<31, 1<<31-1
	for i := 0; i < n; i++ {
		x := pop(i)
		if x > max {
			max = x
		}
		heap.Push(h, p{i, x})
	}

	for {
		x := heap.Pop(h).(p)
		if max-x.v < ans {
			ans = max - x.v
		}
		if len(arr[x.i]) == 0 {
			break
		}

		v := pop(x.i)
		if v > max {
			max = v
		}
		heap.Push(h, p{x.i, v})
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
