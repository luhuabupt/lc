package main

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	vis := make([]bool, n)
	vis[0], vis[firstPerson] = true, true
	tm, ti := map[int]bool{}, []int{}
	for _, v := range meetings {
		if !tm[v[2]] {
			ti = append(ti, v[2])
			tm[v[2]] = true
		}
	}
	sort.Ints(ti)

	tp, tsub := map[int]map[int]int{}, map[int][][]int{}
	for _, v := range meetings {
		if tp[v[2]] == nil {
			tp[v[2]] = map[int]int{}
		}
		if tsub[v[2]] == nil {
			tsub[v[2]] = [][]int{}
		}
		if _, ok := tp[v[2]][v[0]]; !ok {
			if _, ok := tp[v[2]][v[1]]; !ok {
				tp[v[2]][v[0]] = len(tsub[v[2]])
				tp[v[2]][v[1]] = len(tsub[v[2]])

				tsub[v[2]] = append(tsub[v[2]], []int{v[0], v[1]})
			} else {
				tp[v[2]][v[0]] = tp[v[2]][v[1]]
				tsub[v[2]][tp[v[2]][v[1]]] = append(tsub[v[2]][tp[v[2]][v[1]]], v[0])
			}
		} else {
			if _, ok := tp[v[2]][v[1]]; !ok {
				tp[v[2]][v[1]] = tp[v[2]][v[0]]
				tsub[v[2]][tp[v[2]][v[0]]] = append(tsub[v[2]][tp[v[2]][v[0]]], v[1])
			} else {
				np, orip := tp[v[2]][v[0]], tp[v[2]][v[1]]
				if len(tsub[v[2]][orip]) > len(tsub[v[2]][np]) {
					orip, np = np, orip
				}
				for _, x := range tsub[v[2]][orip] {
					tsub[v[2]][np] = append(tsub[v[2]][np], x)
					tp[v[2]][x] = np
				}
				tsub[v[2]][orip] = []int{}
			}
		}
	}

	for _, t := range ti {
		for _, set := range tsub[t] {
			flag := false
			for _, v := range set {
				if vis[v] {
					flag = true
					break
				}
			}
			if flag {
				for _, v := range set {
					vis[v] = true
				}
			}
		}
	}

	ans := []int{}
	for k, v := range vis {
		if v {
			ans = append(ans, k)
		}
	}

	return ans
}

func minimumDeletions(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	i, vi, j, vj := 0, 1000001, 0, -1000001
	for k, v := range nums {
		if vi > v {
			vi, i = v, k
		}
		if vj < v {
			vj, j = v, k
		}
	}
	if i > j {
		i, j = j, i
	}
	//fmt.Println(i, j)

	l, r := 0, len(nums)-1
	x := []int{}
	x = append(x, j-l+1)
	x = append(x, r-i+1)
	x = append(x, i-l+r-j+2)
	sort.Ints(x)
	return x[0]
}
