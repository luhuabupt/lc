package main

import "fmt"

func main() {
	//fmt.Println(groupStrings([]string{"a","b","ab","cde"}))
	fmt.Println(groupStrings([]string{"wrc", "aejd", "x", "w"}))
	//fmt.Println(groupStrings([]string{"a", "b"}))
	//fmt.Println(groupStrings([]string{"ghnv", "uip", "tenv", "hvepx", "e", "ktc", "byjdt", "ulm", "cae", "ea"}))
	//fmt.Println(groupStrings([]string{"k", "et", "zful", "ui", "bz", "nsprc", "kzlj", "aifq"}))
}

func groupStrings(str []string) []int {
	n := len(str)

	// 转二进制
	toBin := func(s string) (x int) {
		for _, v := range s {
			x += 1 << (v - 'a')
		}
		return
	}
	arr := []int{}
	for _, s := range str {
		arr = append(arr, toBin(s))
	}

	m := map[int][]int{}   // 普通字符
	del := map[int][]int{} // 删除字符

	fa := make([]int, n)
	sub := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		sub[i] = []int{i}
	}

	var find func(i int) int
	find = func(i int) int {
		if fa[i] == i {
			return i
		}
		return find(fa[i])
	}

	merge := func(i, j int) {
		a, b := find(i), find(j) // a
		if a == b {
			return
		}
		if len(sub[a]) < len(sub[b]) {
			i, j, a, b = j, i, b, a
		}

		for _, v := range sub[b] {
			fa[v] = a
			sub[a] = append(sub[a], v)
		}
		sub[b] = []int{}
	}

	for i, val := range arr {
		m[val] = append(m[val], i)
		for j := 0; j < 26; j++ {
			if val&(1<<j) == 0 {
				continue
			}
			m[val^(1<<j)] = append(m[val^(1<<j)], i)
			del[val^(1<<j)] = append(m[val^(1<<j)], i)
		}
	}

	// 并查集合并
	for _, arr := range m {
		if len(arr) == 1 {
			continue
		}
		for i := 1; i < len(arr); i++ {
			merge(arr[0], arr[i])
		}
	}
	for _, arr := range del {
		if len(arr) == 1 {
			continue
		}
		for i := 1; i < len(arr); i++ {
			merge(arr[0], arr[i])
		}
	}

	// 计算ans
	ans := []int{0, 0}
	for _, v := range sub {
		if len(v) > 0 {
			ans[0]++
		}
		if len(v) > ans[1] {
			ans[1] = len(v)
		}
	}

	return ans
}
