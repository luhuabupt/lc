package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(123)
}

func analysisHistogram(heights []int, cnt int) []int {
	sort.Ints(heights)
	n := len(heights)
	min := heights[cnt-1] - heights[0]
	minI := 0
	for i := 0; i <= n-cnt; i++ {
		if heights[i+cnt-1]-heights[i] < min {
			min = heights[i+cnt-1] - heights[i]
			minI = i
		}
	}
	return heights[minI : minI+cnt]
}

func metroRouteDesignI(lines [][]int, start int, end int) []int {
	m := map[int]map[int]bool{} // 每个点有几号线
	for i, arr := range lines {
		for _, v := range arr {
			if m[v] == nil {
				m[v] = map[int]bool{}
			}
			m[v][i] = true
		}
	}
	vis := map[int]bool{} // 几号线
	st := []int{}         // 几号线
	for k, _ := range m[start] {
		vis[k] = true
		st = append(st, k)
	}

	find := false
	for !find {
		ns := []int{}
		for _, v := range st {
			for _, nodes := range lines[v] {

				for k, _ := range m[nodes] {
					if m[end][k] {
						find = true
					}
					if !vis[k] {
						ns = append(ns, k)
					}
				}
			}
		}
	}
	return []int{}
}
