package main

import "fmt"

func main() {
	//i := 9
	//if !chk(i%6, (i/6)%6) && !chk((i/6)%6, (i/36)%6) {
	//	fmt.Println(i, i%6+1, (i/6)%6+1, (i/36)%6+1, "do \n")
	//}
	//
	//fmt.Println(!chk(1,3) && !chk(0, 1))
	//fmt.Println(distinctSequences(3))
	fmt.Println(distinctSequences(4)) // 184
	//fmt.Println(distinctSequences(6)) // 1472
}

func chk(a, b int) bool {
	if a == b {
		return true
	}

	if a > b {
		return chk(b, a)
	}
	a++
	b++

	if a == 2 && (b == 4 || b == 6) {
		return true
	}
	if a == 4 && b == 6 {
		return true
	}
	if a == 3 && b == 6 {
		return true
	}

	return false
}

func distinctSequences(n int) int {
	M := int(1e9) + 7
	if n == 1 {
		return 6
	} else if n == 2 {
		return 22
	}
	var chk func(a, b int) bool
	chk = func(a, b int) bool {
		if a == b {
			return true
		}

		if a > b {
			return chk(b, a)
		}
		a++
		b++

		if a == 2 && (b == 4 || b == 6) {
			return true
		}
		if a == 4 && b == 6 {
			return true
		}
		if a == 3 && b == 6 {
			return true
		}

		return false
	}

	p := [36 * 6]int{}
	d := [36 * 6]int{}
	or := []int{}
	for i := 0; i < 36*6; i++ {
		//fmt.Println(i, i%6+1, (i/6)%6+1, (i/36)%6+1)
		if i%6 == (i/6)%6 || i%6 == (i/36)%6 {
			continue
		}
		if !chk(i%6, (i/6)%6) && !chk((i/6)%6, (i/36)%6) {
			//fmt.Println(i, i%6+1, (i/6)%6+1, (i/36)%6+1, "do \n")
			p[i] = 1
			d[i] = 1
			or = append(or, i)
		}
	}
	for i := 4; i <= n; i++ {
		d = [36 * 6]int{}
		for _, j := range or {
			for k := 0; k < 6; k++ {
				if k == j%6 {
					continue
				}
				x := j/6 + k*36
				d[j] += p[x]
				d[j] %= M
				if p[x] > 0 {
					fmt.Println(j, xx(j), x, xx(x), p[x])
				}
			}
		}
		for i := 0; i < 36*6; i++ {
			p[i] = d[i]
		}
	}
	ans := 0
	for _, v := range d {
		ans += v
		ans %= M
	}
	return ans
}

func xx(a int) []int {
	return []int{a/36%6 + 1, a/6%6 + 1, a%6 + 1}
}

func countAsterisks(s string) int {
	ans := 0
	a := 0
	for _, v := range s {
		if v == '|' {
			a++
			a %= 2
		}
		if v == '*' && a == 0 {
			ans++
		}
	}
	return ans
}

func maximumXOR(ar []int) int {
	a := [30]int{}
	for _, v := range ar {
		for i := 0; i < 30; i++ {
			if (1<<i)&v > 0 {
				a[i]++
			}
		}
	}
	ans := 0
	for i, v := range a {
		if v > 0 {
			ans += (1 << i)
		}
	}
	return ans
}

type unionFind struct {
	size int
	fa   []int
	son  [][]int
}

func newUnionFind(n int) *unionFind {
	fa := make([]int, n)
	son := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		son[i] = append(son[i], i)
	}
	return &unionFind{n, fa, son}
}

func (t *unionFind) find(i int) int {
	if t.fa[i] == i {
		return i
	}
	return t.find(t.fa[i])
}

func (t *unionFind) union(i, j int) {
	fi, fj := t.find(i), t.find(j)
	if fi == fj {
		return
	}

	if len(t.son[fi]) < len(t.son[fj]) {
		fi, fj = fj, fi
	}

	for _, v := range t.son[fj] {
		t.fa[v] = fi
	}

	t.son[fi] = append(t.son[fi], t.son[fj]...)
	t.son[fj] = []int{}
	t.size--
}

func countPairs(n int, e [][]int) int64 {
	nx := make([][]int, n)
	for _, v := range e {
		nx[v[0]] = append(nx[v[0]], v[1])
		nx[v[1]] = append(nx[v[1]], v[0])
	}

	uf := newUnionFind(n)
	for i := 0; i < n; i++ {
		for _, x := range nx[i] {
			uf.union(i, x)
		}
	}
	ans := 0
	for _, v := range uf.son {
		ans += len(v) * (n - len(v))
	}
	return int64(ans)
}
