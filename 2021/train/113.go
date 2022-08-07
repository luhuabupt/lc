package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(largestComponentSize([]int{83, 99, 39, 11, 19, 30, 31}))
}

func largestTimeFromDigits(arr []int) string {
	m := map[int]int{}
	for i := 23; i >= 0; i-- {
		si := strconv.Itoa(i)
		if len(si) == 1 {
			si = "0" + si
		}
		for j := 59; j >= 0; j-- {
			ji := strconv.Itoa(j)
			if len(ji) == 1 {
				ji = "0" + ji
			}
			flag := true
			m = map[int]int{}
			for _, v := range arr {
				m[v]++
			}
			for _, v := range si {
				if m[int(v)-'0'] == 0 {
					flag = false
					break
				}
				m[int(v)-'0']--
			}
			for _, v := range ji {
				if m[int(v)-'0'] == 0 {
					flag = false
					break
				}
				m[int(v)-'0']--
			}
			if flag {
				return si + ":" + ji
			}
		}
	}
	return ""
}

// [17,13,11,2,3,5,7]
// [2,13,3,11,5,17,7]
// [2,11,3,17,5,13,7]
func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	sort.Ints(deck)

	ans := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		stack = append(stack, i)
	}

	idx := 0
	for i := 0; i < len(stack); i++ {
		if i%2 == 0 {
			ans[stack[i]] = deck[idx]
			idx++
		} else {
			stack = append(stack, stack[i])
		}
	}
	return ans
}

func largestComponentSize(nums []int) int {
	sort.Ints(nums)
	father := map[int]int{}       // 节点root
	son := map[int]map[int]bool{} // root 子节点
	set := map[int]int{}          // 并查集
	fac := []int{}                // 因数

	for _, v := range nums {
		set[v] = 1
		father[v] = v
		if son[v] == nil {
			son[v] = map[int]bool{}
		}
		son[v][v] = true

		fac = []int{}
		for i := 2; i*i <= v; i++ {
			if v%i == 0 {
				fac = append(fac, i)
				if i != v/i {
					fac = append(fac, v/i)
				}
			}
		}

		if len(fac) == 0 {
			continue
		}

		for _, x := range fac {
			if father[x] != v { // 合并 改father 减set 删son
				oldF := father[x]
				set[v] += set[father[x]]

				set[father[x]] = 0
				father[x] = v
				son[v][x] = true

				for s, _ := range son[oldF] {
					father[s] = v
					son[v][s] = true
				}
				delete(son, oldF)
			}
		}

		//fmt.Println(v, father, son, set)
	}
	ans := 0
	for _, v := range set {
		if v > ans {
			ans = v
		}
	}
	return ans
}
