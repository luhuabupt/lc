package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hs struct{ sort.IntSlice }

func (h *hs) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hs) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() {
	p := [][]int{[]int{3, 10}, []int{1, 5}, []int{2, 6}}
	fmt.Println(smallestChair(p, 0))
}

// 优先队列
func smallestChair(times [][]int, targetFriend int) int {
	ans, arrive, leave, m := 0, [][]int{}, [][]int{}, map[int]int{}
	h := &hp{}
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
