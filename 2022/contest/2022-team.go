package main

import "fmt"

func main() {
	fmt.Println(composeCube([][]string{{"000", "110", "000"}, {"110", "011", "000"}, {"110", "011", "110"}, {"000", "010", "111"}, {"011", "111", "011"}, {"011", "010", "000"}}))
}

func composeCube(sp [][]string) bool {
	cnt := 0
	n := len(sp[0])
	for _, ar := range sp {
		for k, v := range ar {
			if v[0] == '1' {
				cnt++
			}
			if v[len(v)-1] == '1' {
				cnt++
			}
			for i := 1; i < len(v)-1; i++ {
				if v[i] == '1' {
					cnt++
				}
				if k != 0 && k != n-1 && v[i] != '1' {
					return false
				}
			}
		}
	}
	if cnt != n*n*6-12*n+8 {
		return false
	}

	per := permute([]int{0, 1, 2, 3, 4, 5})

	for i := 0; i < 1<<12; i++ {
		for _, p := range per {
			nr := [][][]byte{}
			for j, v := range p {
				do := (i >> (2 * j) & 1) + (i>>(2*j+1)&1)*2
				nr = append(nr, re(sp[v], do))
			}

			if check(nr) {
				return true
			}
		}
	}

	return false
}

func check(sp [][][]byte) bool {
	type D struct {
		i, j   int
		id, jd int // 上下左右
	}
	n := len(sp[0])

	// 12 条棱
	dir := [12]D{
		{0, 1, 2, 0},
		{0, 2, 0, 0},
		{0, 3, 2, 0},
		{0, 4, 1, 0},
		{1, 2, 3, 2},
		{1, 4, 2, 3},
		{1, 5, 1, 3},
		{2, 3, 3, 2},
		{2, 5, 1, 0},
		{3, 4, 3, 2},
		{3, 5, 1, 2},
		{4, 5, 1, 1},
	}

	for _, d := range dir {
		a, b := sp[d.i], sp[d.j]
		x, y := d.id, d.jd
		as, bs := getStr(a, x), getStr(b, y)
		for i := 1; i < n-1; i++ {
			if as[i] == bs[i] {
				return false
			}
		}
	}

	// 8 顶点
	di := [][]int{
		{0, 0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 2, 2, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 2, 0, 0},
		{0, 0, 2, 2, 0, 0, 1, 0, 0},
		{1, 0, 0, 3, 0, 0, 1, 0, 0},
		{1, 0, 3, 2, 0, 2, 1, 0, 0},
		{2, 0, 3, 3, 0, 2, 1, 0, 0},
		{3, 0, 0, 3, 0, 0, 1, 0, 0},
	}
	for _, v := range di {
		xx := []byte{}
		for i := 0; i < 9; i += 3 {
			xx = append(xx, g(sp[v[i]], v[i]+1, v[i]+2))
		}
		cnx := 0
		for _, xxx := range xx {
			if xxx == '1' {
				cnx++
			}
		}
		if cnx != 1 {
			return false
		}
	}

	return true
}

func g(v [][]byte, a, b int) byte {
	n := len(v)
	if a == 0 && b == 2 {
		return v[0][0]
	}
	if a == 0 && b == 3 {
		return v[0][n-1]
	}
	if a == 1 && b == 2 {
		return v[n-1][0]
	}
	return v[n-1][n-1]
}

func getStr(s [][]byte, do int) []byte {
	if do == 0 {
		return s[0]
	}
	if do == 1 {
		return s[len(s)-1]
	}

	n := len(s)
	ans := []byte{}
	if do == 2 {
		for i := 0; i < n; i++ {
			ans = append(ans, s[i][0])
		}
	}
	if do == 3 {
		for i := 0; i < n; i++ {
			ans = append(ans, s[i][n-1])
		}
	}

	return ans
}

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}

func re(v []string, do int) [][]byte {
	ll := len(v)
	nr := make([][]byte, ll)
	if do == 0 {
		for i := 0; i < ll; i++ {
			nr[i] = []byte(v[i])
		}
		return nr
	}

	if do == 1 {
		for i := 0; i < ll; i++ {
			nr[i] = make([]byte, ll)
			for j := 0; j < ll; j++ {
				nr[i][j] = v[ll-1-j][i]
			}
		}
		return nr
	}

	if do == 2 {
		for i := 0; i < ll; i++ {
			nr[i] = make([]byte, ll)
			for j := 0; j < ll; j++ {
				nr[i][j] = v[ll-1-i][ll-1-j]
			}
		}
		return nr
	}

	if do == 3 {
		for i := 0; i < ll; i++ {
			nr[i] = make([]byte, ll)
			for j := 0; j < ll; j++ {
				nr[i][j] = v[j][ll-1-i]
			}
		}
		return nr
	}

	return nr
}
