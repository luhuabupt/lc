package main

import "fmt"

func main() {
	fmt.Println(" ")
}

type tree struct {
	arr []int
}

func (t tree) add(i, v int) {
	for i < len(t.arr) {
		t.arr[i] = t.arr[i] + v
		i += lowbit(i)
	}
}

func (t tree) sum(i int) (ans int) {
	for i > 0 {
		ans += t.arr[i]
		i -= lowbit(i)
	}
	return
}

func (t tree) query(l, r int) (ans int) {
	return t.sum(r) - t.sum(l-1)
}

func lowbit(v int) int {
	return v & -v
}

func numberOfPairs(a []int, b []int, d int) int64 {
	n := len(a)
	t := tree{make([]int, 3*1e5)}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		t.add(d-b[i]+a[i]+1e5, 1)
		ans += t.query(a[i]-b[i]+1e5, 3*1e5)
	}
	return int64(ans)
}

type LUPrefix struct {
	fa   []int
	fa0  int
	son  [][]int
	size []int
}

func (t *LUPrefix) find(i int) int {
	if t.fa[i] == i {
		return i
	}
	return t.find(t.fa[i])
}

func (t *LUPrefix) union(i, j int) {
	fi, fj := t.find(i), t.find(j)
	if fi == fj {
		return
	}
	if t.size[fi] < t.size[fj] {
		fi, fj = fj, fi
	}

	aj := t.son[fj]
	t.fa[fj] = fi
	if t.fa0 == fj {
		t.fa0 = fi
	}
	if aj[0] < t.son[fi][0] {
		t.son[fi][0] = aj[0]
	}
	if aj[1] > t.son[fi][1] {
		t.son[fi][1] = aj[1]
	}
}

func Constructor(n int) LUPrefix {
	fa := make([]int, n)
	son := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		son[i] = []int{-1, -1}
	}
	return LUPrefix{fa, 0, son, make([]int, n)}
}

func (t *LUPrefix) Upload(v int) {
	v--
	t.son[v] = []int{v, v}
	t.size[v] = 1
	if v > 0 && t.size[v-1] > 0 {
		t.union(v, v-1)
	}
	if v < len(t.fa)-1 && t.size[v+1] > 0 {
		t.union(v, v+1)
	}
}

func (t *LUPrefix) Longest() int {
	// fmt.Println(t.fa)
	// fmt.Println(t.son)

	return t.son[t.fa0][1] + 1
}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */

func equalFrequency(a string) bool {
	c := make([]int, 26)
	for _, v := range a {
		c[v-'a']++
	}

	chk := func(v int) bool {
		nc := append([]int{}, c...)
		nc[v]--
		m := map[int]int{}
		for _, x := range nc {
			if x == 0 {
				continue
			}
			m[x]++
		}
		return len(m) == 1
	}

	for _, v := range a {
		if chk(int(v - 'a')) {
			return true
		}
	}
	return false

	// c := [26]int{}
	// for _, v := range a {
	//     c[v-'a']++
	// }
	// b := map[int]int{}
	// for _, v := range c {
	//     if v > 0 {
	//         b[v]++
	//     }
	// }
	// fmt.Println(b)
	// if len(b) == 2 {
	//     x, y := 0, 0
	//     for k, _ := range b {
	//         if x == 0 {
	//             x = k
	//         } else if y == 0 {
	//             y = k
	//         }
	//     }
	//     if x== 1 || y == 1 {
	//         return true
	//     }
	//     if x > y {
	//         x, y = y, x
	//     }
	//     if b[y] != 1 {
	//         return false
	//     }
	//     return y == x+1 || y == 1
	// }
	// return false
}
