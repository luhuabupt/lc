package main

import "fmt"

func main() {
	a := [][]int{{1, 2}, {2, 1}, {3, 4}}
	fmt.Println(topoSort(a, []int{1, 2, 4}))
}

// 拓扑排序
func topoSort(a [][]int, all []int) []int {
	ans := []int{}

	in := map[int]int{} // 入度
	nx := map[int][]int{}
	for _, v := range a {
		in[v[1]]++

		if nx[v[0]] == nil {
			nx[v[0]] = []int{}
		}
		nx[v[0]] = append(nx[v[0]], v[1])
	}

	s := []int{} // 入度为0
	for _, v := range all {
		if in[v] == 0 {
			s = append(s, v)
		}
	}

	for len(s) > 0 {
		ns := []int{}
		for _, x := range s {
			ans = append(ans, x)
			for _, v := range nx[x] { // 拆边
				in[v]--
				if in[v] == 0 {
					ns = append(ns, v)
				}
			}
		}
		s = ns
	}

	if len(all) != len(ans) { // 有环
		return []int{}
	}

	return ans
}
