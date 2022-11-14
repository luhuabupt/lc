package main

import "fmt"

func main() {
	fmt.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {2, 3}}, 2, []int{-5644, -6018, 1188, -8502}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	nx := make([][]int, n)
	for _, v := range edges {
		nx[v[0]] = append(nx[v[0]], v[1])
		nx[v[1]] = append(nx[v[1]], v[0])
	}

	vis := make([]bool, n)
	db := make([]int, n)
	for i := 0; i < n; i++ {
		db[i] = -1
	}
	var d0 func(i, d int) int
	d0 = func(i, d int) int {
		if i == 0 {
			db[i] = d
			return d - 1
		}
		vis[i] = true
		r := -1
		for _, x := range nx[i] {
			if vis[x] {
				continue
			}
			r = d0(x, d+1)
			if r >= 0 {
				break
			}
		}
		db[i] = r
		return r - 1
	}
	d0(bob, 0)
	//fmt.Println(db)
	vis = make([]bool, n)
	ans := -1 << 62
	var dfs func(i, d, s int)
	dfs = func(i, d, s int) {
		vis[i] = true
		if db[i] < 0 || d < db[i] {
			s += amount[i]
		} else if d == db[i] {
			s += amount[i] / 2
		}

		isLeaf := true
		for _, x := range nx[i] {
			if vis[x] {
				continue
			}
			isLeaf = false
			dfs(x, d+1, s)
		}

		//fmt.Println(i,d,s,isLeaf)
		if isLeaf && s > ans {
			ans = s
		}
	}
	dfs(0, 0, 0)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
