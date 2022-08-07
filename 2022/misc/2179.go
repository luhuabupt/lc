package main

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

func goodTriplets(a []int, b []int) int64 {
	mb := map[int]int{}
	for k, v := range b {
		mb[v] = k
	}

	ans := int64(0)
	t := tree{make([]int, len(a)+1)}

	for k, v := range a {
		x := t.query(1, mb[v])
		ans += int64(x) * int64(len(a)-1-mb[v]-k+x)

		t.add(mb[v]+1, 1)
	}

	return ans
}
